package gh

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	ut "utils"

	"github.com/google/go-github/github"
)

/*
GitClient g
*/
type GitClient struct {
	Token string
}

func getGitClient2() *GitClient {
	token := GetToken()

	if token == nil {
		return &GitClient{" "}
	}

	return &GitClient{token.Token}
}

/*
HTTPGet http
*/
func (client *GitClient) HTTPGet(url string, headers map[string]string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	ut.HandleError(err)
	return client.HTTPRequest(req, headers)
}

/*
HTTPPost http
*/
func (client *GitClient) HTTPPost(url string, body string, headers map[string]string) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(body))
	ut.HandleError(err)
	return client.HTTPRequest(req, headers)
}

/*
HTTPRequest g
*/
func (client *GitClient) HTTPRequest(req *http.Request, headers map[string]string) *http.Response {

	// req.Header.Add("Content-Type", "application/vnd.github.v3+json")
	req.Header.Add("Content-Type", "application/json")
	// req.Header.Add("userAgent", "x-cmd")

	// req.Header.Add("Accept", "application/json")

	req.Header.Add("Authorization", "token "+client.Token)

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

/*
EnablePage a
*/
func (client *GitClient) EnablePage(owner, repo, branch, folder string) error {

	headers := map[string]string{}

	// https://developer.github.com/changes/2019-03-14-enabling-disabling-pages/
	var mediaTypeEnablePagesAPIPreview = "application/vnd.github.switcheroo-preview+json"

	// https://developer.github.com/changes/2016-07-06-github-pages-preiew-api/
	var mediaTypePagesPreview = "application/vnd.github.mister-fantastic-preview+json"

	acceptHeaders := []string{mediaTypeEnablePagesAPIPreview, mediaTypePagesPreview}
	headers["Accept"] = strings.Join(acceptHeaders, ", ")

	var strMaster = branch // "master"
	var strPath = folder   // "/docs"

	var ghAPI = "https://api.github.com/"

	resp := client.HTTPPost(
		fmt.Sprintf(ghAPI+"repos/%v/%v/pages", owner, repo), ut.JS(&github.Pages{
			Source: &github.PagesSource{
				Branch: &strMaster,
				Path:   &strPath,
			},
		}), headers)

	b, _ := ioutil.ReadAll(resp.Body)
	// log.WithFields(map[string]interface{}{
	// 	"status": resp.Status,
	// 	"body":   string(b),
	// }).Debug("EnablePages result")

	if resp.StatusCode == 201 {
		return nil
	}
	return errors.New(string(b))

	// p, _, e1 := getClient().Repositories.EnablePages(getCtx(), owner, repo)
	// fmt.Println(ut.JS(p))
	// return e1

	// getClient().Repositories.
}
