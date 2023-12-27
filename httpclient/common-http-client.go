package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func ExecuteApi(method string, url string, body string, headers map[string]string) ([]byte, int, error) {
	req, err := GetRequest(method, url, body)
	if err != nil {
		log.Println(err.Error())
		return nil, http.StatusInternalServerError, err
	}
	if headers == nil {
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
	} else {
		requestHeaders := make(http.Header)
		for key, value := range headers {
			requestHeaders.Set(key, value)
		}
		req.Header = requestHeaders
	}

	//req.Header.Add("Accept", "application/json")
	//req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil && res != nil {
		fmt.Println("error in getting response from server", "url", url, "method", req.Method, "error", err.Error(), "status code", res.StatusCode)
		return nil, res.StatusCode, fmt.Errorf("error in getting response from server. api url %s", url)
	}
	if err != nil && res == nil {
		fmt.Println("error getting response from server. no response received", "url", url, "error", err.Error())
		return nil, http.StatusInternalServerError, fmt.Errorf("error getting response from server. no response received. api url %s. no response received. Error: %s", url, err.Error())
	}
	if err == nil && res == nil {
		fmt.Println("invalid response from server and also no error", "url", url, "method", req.Method)
		return nil, http.StatusInternalServerError, fmt.Errorf("invalid response from server and also no error. api url %s", url)
	}
	if res.StatusCode >= http.StatusBadRequest {
		return nil, res.StatusCode, fmt.Errorf(res.Status)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response body", "url", url, "error", err.Error())
		return nil, res.StatusCode, err
	}
	bodyBytes = RemoveBOMContent(bodyBytes)
	return bodyBytes, res.StatusCode, err
}

func RemoveBOMContent(input []byte) []byte {
	return bytes.TrimPrefix(input, []byte("\xef\xbb\xbf"))
}

func GetRequest(method string, url string, body string) (req *http.Request, err error) {
	switch strings.ToUpper(method) {
	case http.MethodPost:
		var bd io.Reader
		if body != "" {
			bd = strings.NewReader(body)
		}
		req, err = http.NewRequest(http.MethodPost, url, bd)
	default:
		req, err = http.NewRequest(http.MethodGet, url, nil)
	}
	return req, err
}
