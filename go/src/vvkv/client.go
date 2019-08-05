package vvkv

import (
	"encoding/json"
	"fluent/list"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
	ut "utils"
)

func checkStatusCode(resp *http.Response) bool {
	switch resp.StatusCode {
	case 200:
		return true
	case 403:
		panic("Access Denied")
	case 404:
		// Log.Info("URL Not Found: " + resp.Request.URL.String())
		return false
	default:
		panic("Status code NOT 200: " + resp.Request.URL.String())
	}
}

/*
DecryptToken a
*/
func (client *Client) DecryptToken(token string) string {
	url := client.URL + "/token-decrypt"
	res := client.HTTPPost(url, token, nil)
	body, err := ioutil.ReadAll(res.Body)
	ut.HandleError(err)
	return string(body)
}

/*
GetVVKVURIToLocalDisk Download file to destination path
*/
func (client *Client) GetVVKVURIToLocalDisk(vvkvURI string, dst string, dstForIdx string) *Meta {
	url := client.URL + "/kv-get/" + vvkvURI[1:]
	return client.getToLocalDisk(url, dst, dstForIdx)
}

/*
GetToLocalDisk Download file to destination path
*/
func (client *Client) GetToLocalDisk(org string, key string, dst string, dstForIdx string) *Meta {
	url := client.URL + "/kv-get/" + org + "/" + key
	return client.getToLocalDisk(url, dst, dstForIdx)
}

func (client *Client) getToLocalDisk(url string, dst string, dstForIdx string) *Meta {
	resp := client.HTTPGet(url, nil)
	defer resp.Body.Close()
	ok := checkStatusCode(resp)
	if !ok {
		return nil
	}

	// TODO: Should use buffer tech to avoid writing disk
	temppath, temppathErr := ioutil.TempFile(os.TempDir(), "vvsh*.zip")
	ut.HandleError(temppathErr)
	_, err := io.Copy(temppath, resp.Body)
	ut.HandleError(err)

	ut.Mkdirp(path.Dir(dst))
	UnzipOneFile(temppath.Name(), dst)
	defer os.Remove(temppath.Name())

	meta := resp.Header.Get("x-vvkv-m")
	ut.Mkdirp(path.Dir(dstForIdx))
	ut.WriteFileSync(dstForIdx, meta)

	var ret Meta
	json.Unmarshal([]byte(meta), &ret)

	return &ret
}

/*
ListVVURL all files according to prefix
*/
func (client *Client) ListVVURL(vvurl string) []ListResultItem {
	url := client.URL + "/kv-list/" + vvurl[1:]
	return client.list(url)
}

/*
List all files according to prefix
*/
func (client *Client) List(org string, prefix string) []ListResultItem {
	url := client.URL + "/kv-list/" + org + "/" + prefix
	return client.list(url)
}

func (client *Client) list(url string) []ListResultItem {
	resp := client.HTTPGet(url, nil)
	defer resp.Body.Close()
	checkStatusCode(resp)

	respBody, _ := ioutil.ReadAll(resp.Body)
	respBodyStr := string(respBody)
	var c []ListResultItem
	json.Unmarshal([]byte(respBodyStr), &c)
	return c
}

/*
UploadByVVURL file to key
*/
func (client *Client) UploadByVVURL(src string, vvurl string, isPublic bool, meta map[string]string) {
	url := client.URL + "/kv-put/" + vvurl[1:]
	client.upload(src, url, isPublic, meta)
}

/*
Upload file to key
*/
func (client *Client) Upload(src string, org string, key string, isPublic bool, meta map[string]string) {
	url := client.URL + "/kv-put/" + org + "/" + key
	client.upload(src, url, isPublic, meta)
}

func (client *Client) upload(src string, url string, isPublic bool, meta map[string]string) {
	access := "private"
	if isPublic {
		access = "public"
	}
	meta["x-vvkv-access"] = access

	temppath, temppathErr := ioutil.TempFile(os.TempDir(), "vvsh*.zip")
	ut.HandleError(temppathErr)
	ZipFiles(temppath.Name(), []string{src})
	defer os.Remove(temppath.Name())

	// TODO: hidden bug to resolve. Zip should be a binary file. should use byte[]. Not string.
	content := ut.ReadFile(temppath.Name())
	resp := client.HTTPPost(url, content, meta)
	defer resp.Body.Close()
	checkStatusCode(resp)
}

/*
SetPermissionBYVVURL file to public or private
*/
func (client *Client) SetPermissionBYVVURL(vvurl string, isPublic bool) {
	url := client.URL + "/kv-access/" + vvurl[1:]
	client.setPermission(url, isPublic)
}

/*
SetPermission file to public or private
*/
func (client *Client) SetPermission(org string, key string, isPublic bool) {
	url := client.URL + "/kv-access/" + org + "/" + key
	client.setPermission(url, isPublic)
}

func (client *Client) setPermission(url string, isPublic bool) {

	access := "private"
	if isPublic {
		access = "public"
	}

	resp := client.HTTPGet(url, map[string]string{
		"x-vvkv-access": access,
	})
	defer resp.Body.Close()
	checkStatusCode(resp)
}

/*
Delete file to public or private
TODO: One day, should use delete method
*/
func (client *Client) Delete(org string, key string) {
	url := client.URL + "/kv-delete/" + org + "/" + key
	resp := client.HTTPGet(url, nil)
	defer resp.Body.Close()
	checkStatusCode(resp)
}

type tokenApplication struct {
	Get     []string `json:"get"`
	Put     []string `json:"put"`
	Info    string   `json:"info"`
	Expired int64    `json:"expired"`
}

/*
ApplyToken Share generate shorturl for
*/
func (client *Client) ApplyToken(get []string, put []string, info string, timestamp int64) string {
	url := client.URL + "/token-encrypt"
	st := ut.JS(&tokenApplication{get, put, info, timestamp})
	resp := client.HTTPPost(url, st, nil)

	ret, err := ioutil.ReadAll(resp.Body)
	ut.HandleError(err)
	return string(ret)
}

/*
Share generate shorturl for
*/
func (client *Client) Share(vvurl string, shareTime int64) string {
	token := client.ApplyToken(list.Str(vvurl[1:]), list.Str(), "share", time.Now().Unix()+shareTime)
	return client.URL + "/kv-cat/" + vvurl[1:] + "?token=" + token
}
