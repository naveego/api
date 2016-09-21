package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	basePath          string
	version           string
	customHTTPHeaders map[string]string
	httpClient        *http.Client
}

func NewClient(host, version string, httpHeaders map[string]string) (*Client, error) {

	httpClient := &http.Client{}

	return &Client{
		basePath:          host,
		version:           version,
		httpClient:        httpClient,
		customHTTPHeaders: httpHeaders,
	}, nil
}

type serverResponse struct {
	body       io.Reader
	header     http.Header
	statusCode int
}

func (cli *Client) get(path string, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest("GET", path, nil, headers)
}

func (cli *Client) post(path string, data interface{}, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest("POST", path, data, headers)
}

func (cli *Client) put(path string, data interface{}, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest("PUT", path, data, headers)
}

func (cli *Client) delete(path string, headers map[string][]string) (serverResponse, error) {
	return cli.sendRequest("DELETE", path, nil, headers)
}

func (cli *Client) sendRequest(method, path string, obj interface{}, headers map[string][]string) (serverResponse, error) {
	var body io.Reader

	if obj != nil {
		var err error
		body, err = encodeData(obj)
		if err != nil {
			return serverResponse{}, err
		}
		if headers == nil {
			headers = make(map[string][]string)
		}
		headers["Content-Type"] = []string{"application/json"}
	}

	serverResp := serverResponse{
		body:       nil,
		statusCode: -1,
	}

	expectedPayload := (method == "POST" || method == "PUT")
	if expectedPayload && body == nil {
		body = bytes.NewReader([]byte{})
	}

	req, err := cli.newRequest(method, path, body, headers)
	if err != nil {
		return serverResp, err
	}

	if expectedPayload && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "text/plain")
	}

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return serverResp, fmt.Errorf("An error occurred trying to connect: %v", err)
	}
	defer resp.Body.Close()

	if resp != nil {
		serverResp.statusCode = resp.StatusCode
	}

	if serverResp.statusCode < 200 || serverResp.statusCode >= 400 {
		return serverResp, fmt.Errorf("Error response from server: %s", "")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return serverResp, err
	}

	serverResp.body = ioutil.NopCloser(bytes.NewReader(respBody))
	serverResp.header = resp.Header
	return serverResp, nil
}

func (cli *Client) newRequest(method, path string, body io.Reader, headers map[string][]string) (*http.Request, error) {
	apiPath := cli.basePath + path
	req, err := http.NewRequest(method, apiPath, body)
	if err != nil {
		return nil, err
	}

	for k, v := range cli.customHTTPHeaders {
		req.Header.Set(k, v)
	}

	if headers != nil {
		for k, v := range headers {
			req.Header[k] = v
		}
	}

	return req, nil
}

func encodeData(data interface{}) (*bytes.Buffer, error) {
	params := bytes.NewBuffer(nil)
	if data != nil {
		if err := json.NewEncoder(params).Encode(data); err != nil {
			return nil, err
		}
	}
	return params, nil
}
