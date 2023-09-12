package httputils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func queryFromMap(query map[string]string) url.Values {
	q := url.Values{}
	for key, value := range query {
		q.Add(key, value)
	}
	return q
}

func SendPlainRequest(
	endpoint string,
	path string,
	method string,
	query map[string]string,
) (*http.Response, error) {
	client := &http.Client{}
	requestURL, _ := url.ParseRequestURI(endpoint)
	requestURL.Path = path
	q := queryFromMap(query)
	requestURL.RawQuery = q.Encode()

	req, err := http.NewRequest(method, fmt.Sprintf("%v", requestURL), nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func SendRequestWithBody[T any](
	endpoint string,
	path string,
	method string,
	query map[string]string,
	body T,
	headers ...http.Header,
) (*http.Response, error) {
	client := &http.Client{}
	requestURL, _ := url.ParseRequestURI(endpoint)
	requestURL.Path = path
	q := queryFromMap(query)
	requestURL.RawQuery = q.Encode()

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%v", requestURL), &buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if len(headers) > 0 {
		for key, value := range headers[0] {
			req.Header[key] = value
		}
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ReadJSONRes[T any](res *http.Response) (*T, error) {
	var payload T

	err := json.NewDecoder(res.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
