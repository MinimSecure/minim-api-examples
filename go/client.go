package minim

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
)

const oauthTokenPath = "/api/oauth/token"

type Client struct {
	Host string

	applicationID string
	secret        string

	token string

	cl *http.Client
}

func New(appId, secret string) *Client {
	return &Client{
		applicationID: appId,
		secret:        secret,
		Host:          "my.minim.co",
		cl:            &http.Client{},
	}
}

type QueryParams map[string]string
type JSONBody interface{}
type JSONObject map[string]interface{}

func (cl *Client) currentToken() (string, error) {
	if cl.token != "" {
		return cl.token, nil
	}
	data := JSONObject{
		"client_id":     cl.applicationID,
		"client_secret": cl.secret,
		"grant_type":    "client_credentials",
	}
	resp, err := cl.PostJSON(oauthTokenPath, JSONBody(data))
	if err != nil {
		return "", err
	}
	tok, err := parseToken(resp)
	if err != nil {
		return "", err
	}
	cl.token = tok
	return cl.token, nil
}

func parseIDs(resp *http.Response) ([]string, error) {
	recs, err := ParseJSONObjects(resp)
	if err != nil {
		return nil, err
	}
	var res []string
	for _, v := range recs {
		id, ok := v["id"].(string)
		if !ok {
			log.Println("expected string for id but got", v["id"])
			continue
		}
		res = append(res, id)
	}
	return res, nil
}

func (cl *Client) GetIDs(url string) ([]string, error) {
	var ids []string
	offset := 0
	for {
		resp, err := cl.GetParams(url, QueryParams{"offset": strconv.Itoa(offset)})
		if err != nil {
			return nil, err
		}
		h, ok := resp.Header["X-Total-Count"]
		var total int64
		if ok && len(h) > 0 {
			total, err = strconv.ParseInt(h[0], 10, 64)
			if err != nil {
				log.Println("failed to parse X-Total-Count header:", err)
			}
		}
		is, err := parseIDs(resp)
		if err != nil {
			return nil, err
		}
		ids = append(ids, is...)
		offset = len(ids)
		if offset >= int(total) {
			break
		}
	}
	return ids, nil
}

func (cl *Client) MultiGet(url string, ids []string) ([]JSONObject, error) {
	return cl.MultiGetParams(url, ids, nil)
}

func (cl *Client) MultiGetParams(url string, ids []string, params QueryParams) ([]JSONObject, error) {
	if len(ids) == 0 {
		var err error
		ids, err = cl.GetIDs(url)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get IDs")
		}
	}
	var res []JSONObject
	for _, v := range ids {
		qp := QueryParams{"id": v}
		for k, v := range params {
			qp[k] = v
		}
		resp, err := cl.GetParams(url, qp)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to fetch details")
		}
		recs, err := ParseJSONObjects(resp)
		if err != nil {
			return nil, errors.Wrapf(err, "unable to parse json response")
		}
		res = append(res, recs...)
	}
	return res, nil
}

func (cl *Client) Get(url string) (*http.Response, error) {
	return cl.request("GET", url, nil, nil, nil)
}

func (cl *Client) GetParams(url string, params QueryParams) (*http.Response, error) {
	return cl.request("GET", url, params, nil, nil)
}

type readCloser struct {
	io.Reader
}

func (rc *readCloser) Close() error {
	return nil
}

func (cl *Client) PostJSON(url string, body JSONBody) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal request body into JSON")
	}
	r := &readCloser{bytes.NewReader(b)}
	return cl.request("POST", url, nil, r, http.Header{"Content-type": {"application/json"}})
}

func (cl *Client) PatchJSON(url string, body JSONBody) (*http.Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal request body into JSON")
	}
	r := &readCloser{bytes.NewReader(b)}
	return cl.request("PATCH", url, nil, r, http.Header{"Content-type": {"application/json"}})
}
