package vvkv

import (
	"bytes"
	"net/http"
	ut "utils"
)

/*
Client for vvkv
*/
type Client struct {
	Token string
	URL   string
}

/*
CreateClient with token
*/
func CreateClient(url string) *Client {
	return &Client{ReadToken(), url}
}

// func getURL(org string, key string) string {
// 	return vvkv_url + "/" + org + "/" + key
// }

/*
HTTPGet http
*/
func (client *Client) HTTPGet(url string, headers map[string]string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	ut.HandleError(err)
	return client.HTTPRequest(req, headers)
}

/*
HTTPPost http
*/
func (client *Client) HTTPPost(url string, body string, headers map[string]string) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	ut.HandleError(err)
	return client.HTTPRequest(req, headers)
}

/*
HTTPRequest http
*/
func (client *Client) HTTPRequest(req *http.Request, headers map[string]string) *http.Response {

	// req.Header.Add("Accept", "application/json")
	req.Header.Add("x-vvkv-token", client.Token)

	if nil != headers {
		for key, value := range headers {
			req.Header.Add(key, value)
		}
	}

	cli := &http.Client{}
	resp, err := cli.Do(req)
	ut.HandleError(err)
	return resp
}
