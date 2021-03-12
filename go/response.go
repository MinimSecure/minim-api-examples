package minim

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

func CloseResponse(resp *http.Response) {
	if resp.Body != nil {
		if err := resp.Body.Close(); err != nil {
			log.Printf("failed to close response body: %s", err)
		}
	}
}

func ParseJSONBody(resp *http.Response) (JSONBody, error) {
	defer CloseResponse(resp)
	var data JSONBody
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read response body")
	}
	err = json.Unmarshal(b, data)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal response body as JSON")
	}
	return data, nil
}

func ParseJSONObject(resp *http.Response) (JSONObject, error) {
	defer CloseResponse(resp)
	data := JSONObject{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read response body")
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal response body as JSON")
	}
	return data, nil
}

func ParseJSONObjects(resp *http.Response) ([]JSONObject, error) {
	defer CloseResponse(resp)
	var data []JSONObject
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read response body")
	}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal response body as JSON")
	}
	return data, nil
}

func parseError(resp *http.Response) (error, error) {
	defer CloseResponse(resp)
	data := JSONObject{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to read response body")
	}
	log.Println(string(b))
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to unmarshal response body as JSON")
	}
	rt, ok := data["error"]
	if !ok {
		return nil, errors.New("response did not contain error")
	}
	ec, ok := rt.(string)
	if !ok {
		return nil, errors.Errorf("response contained error but it was not a string: %v", rt)
	}
	ed := "(no description)"
	rt, ok = data["error_description"]
	if ok {
		ed, ok = rt.(string)
		if !ok {
			ed = "(unable to parse description)"
		}
	}
	return errors.Errorf("%s: %s", ec, ed), nil
}

func parseToken(resp *http.Response) (string, error) {
	defer CloseResponse(resp)
	data := JSONObject{}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrapf(err, "unable to read response body")
	}
	log.Println(string(b))
	err = json.Unmarshal(b, &data)
	if err != nil {
		return "", errors.Wrapf(err, "unable to unmarshal response body as JSON")
	}
	rt, ok := data["access_token"]
	if !ok {
		return "", errors.New("response did not contain access_token")
	}
	stok, ok := rt.(string)
	if !ok {
		return "", errors.Errorf("response contained access_token but it was not a string: %v", stok)
	}
	return stok, nil
}
