package httpP

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

func getResp(client *http.Client, req *http.Request) (body []byte, err error) {
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return
}

func newGitLabRequest(method string, url string, body any) (req *http.Request, err error) {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewBuffer(bodyBytes)
	}

	req, err = http.NewRequest(method, url, bodyReader)
	if err != nil {
		return
	}
	if strings.ToUpper(method) == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}
	return
}

func GetHttpReqResp(client *http.Client, method string, url string, body any) (resp []byte, err error) {
	req, err := newGitLabRequest(method, url, body)
	if err != nil {
		return
	}
	return getResp(client, req)
}
