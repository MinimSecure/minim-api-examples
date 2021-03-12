package minim

import (
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/pkg/errors"
)

func (cl *Client) buildRequest(method, p string, params QueryParams, data io.ReadCloser, headers http.Header) (*http.Request, error) {
	us := "https://" + path.Join(cl.Host, p)
	u, err := url.Parse(us)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to parse URL: %s", us)
	}
	qv := url.Values{}
	for k, v := range params {
		qv[k] = []string{v}
	}
	u.RawQuery = qv.Encode()
	h := http.Header{
		"User-Agent": {"Minim API Examples (Go)"},
		"Accept":     {"application/json"},
	}
	for k, v := range headers {
		h[k] = append([]string{}, v...)
	}
	req := &http.Request{
		Method: method,
		URL:    u,
		Body:   data,
		Header: headers,
	}
	return req, nil
}

func (cl *Client) request(method, p string, params QueryParams, data io.ReadCloser, headers http.Header) (*http.Response, error) {
	// Do not try to get a token when we are already trying to get one.
	if p != oauthTokenPath {
		tok, err := cl.currentToken()
		if err != nil {
			return nil, errors.Wrapf(err, "unable to get valid token")
		}
		if params == nil {
			params = QueryParams{}
		}
		params["access_token"] = tok
	}
	req, err := cl.buildRequest(method, p, params, data, headers)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to prepare request")
	}
	resp, err := cl.cl.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "request failed")
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		e, err := parseError(resp)
		if err != nil {
			return nil, errors.Wrapf(err, "response status is %d, but unable to parse error in response", resp.StatusCode)
		}
		return nil, errors.Wrapf(e, "response status is %d", resp.StatusCode)
	}
	return resp, nil
}
