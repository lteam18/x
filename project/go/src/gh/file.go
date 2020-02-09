package gh

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"
	ut "utils"

	"github.com/google/go-github/github"
)

/*
File a
*/
type File struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	Keypath string `json:"keypath"`
}

/*
Cat a
*/
func (gf *File) Cat() (string, error) {
	c := GetClient()
	f, _, resp, err := c.Repositories.GetContents(GetCtx(), gf.Owner, gf.Repo, gf.Keypath, nil)
	if err != nil {
		return "", err
	}
	if f != nil {
		return f.GetContent()
	}

	// var sw io.StringWriter
	bte, _ := ioutil.ReadAll(resp.Body)
	return "", errors.New(string(bte))
}

/*
CopyTo a
*/
func (gf *File) CopyTo(dstKP string) error {
	s, e := gf.Cat()
	if e != nil {
		return e
	}
	dstGF := &File{gf.Owner, gf.Repo, dstKP}
	if e := dstGF.Upload([]byte(s)); e != nil {
		return e
	}
	return nil
}

/*
GetGHPageURL g
*/
func (gf *File) GetGHPageURL() string {
	return fmt.Sprintf("https://%v.github.io/%v/%v", gf.Owner, gf.Repo, gf.Keypath)
}

/*
GetGithubPageWithLocalCache s
*/
func (gf *File) DownloadFromGithubPageWithCache(localFilePath string, update bool) error {
	if update == false {
		if ut.IsFileExisted(localFilePath) {
			// log.Debugf("Using Local Cache %s", localFilePath)
			return nil
		}
	}

	publicURL := fmt.Sprintf("https://%v.github.io/%v/%v", gf.Owner, gf.Repo, gf.Keypath)
	// log.WithField("url", publicURL).Debug("download")
	ut.Mkdirp(path.Dir(localFilePath))
	if err := ut.HTTPCat(publicURL, localFilePath); err != nil {
		return err
	}
	return nil
}

/*
DownloadPublicFileWithCache g
*/
// func (gf *GithubFile) DownloadPublicFileWithCache(localFilePath string, update bool) error {

// 	tk := GetToken()
// 	if tk != nil {
// 		if tk.Owner != "" && tk.Owner == gf.Owner {
// 			e1 := gf.DownloadWithLocalCache(localFilePath, update)
// 			if e1 == nil {
// 				return nil
// 			}
// 		}
// 	}

// 	e := gf.DownloadFromGithubPageWithCache(localFilePath, update)
// 	if e != nil {
// 		return e
// 	}

// 	return nil
// }

/*
GetPublicFileContentWithCache g
*/
// func (gf *GithubFile) GetPublicFileContentWithCache(localFilePath string, update bool) (string, error) {
// 	e := gf.DownloadPublicFileWithCache(localFilePath, update)
// 	if e != nil {
// 		return "", e
// 	}
// 	return ut.ReadFile(localFilePath), nil
// }

/*
DownloadWithLocalCache g
*/
func (gf *File) DownloadWithLocalCache(localFilePath string, update bool) error {
	if update == false {
		if ut.IsFileExisted(localFilePath) {
			// log.Debugf("Using Local Cache %s", localFilePath)
			return nil
		}
	}

	content, e := gf.Cat()
	if e != nil {
		return e
	}
	ut.Mkdirp(path.Dir(localFilePath))
	ut.WriteFileSync(localFilePath, content)
	return nil
}

/*
RetrieveWithLocalCache g
*/
func (gf *File) RetrieveWithLocalCache(localFilePath string, update bool) (string, error) {
	if update == false {
		if ut.IsFileExisted(localFilePath) {
			// log.Debugf("Using Local Cache %s", localFilePath)
			return ut.ReadFile(localFilePath), nil
		}
	}

	content, e := gf.Cat()
	if e != nil {
		return "", e
	}
	ut.Mkdirp(path.Dir(localFilePath))
	ut.WriteFileSync(localFilePath, content)
	return content, nil
}

/*
GetTempURL g
*/
func (gf *File) GetTempURL() (string, error) {
	c := GetClient()
	f, _, resp, err := c.Repositories.GetContents(GetCtx(), gf.Owner, gf.Repo, gf.Keypath, nil)
	if err != nil {
		return "", err
	}
	if f != nil {
		return f.GetDownloadURL(), nil
	}

	// var sw io.StringWriter
	bte, _ := ioutil.ReadAll(resp.Body)
	return "", errors.New(string(bte))
}

/*
UploadFile a
*/
func (gf *File) Upload(content []byte) error {
	c := GetClient()
	msg := "upload"
	// sha := ut.CalSHA1(string(content))
	fileContent, _, _, getErr := c.Repositories.GetContents(GetCtx(), gf.Owner, gf.Repo, gf.Keypath, nil)

	SHA := ""
	if getErr == nil {
		SHA = *fileContent.SHA
	}
	_, _, err := c.Repositories.CreateFile(GetCtx(), gf.Owner, gf.Repo, gf.Keypath,
		&github.RepositoryContentFileOptions{
			Message:   &msg,
			Content:   content,
			SHA:       &SHA,
			Branch:    nil,
			Author:    nil,
			Committer: nil,
		})
	return err
}

/*
Delete a
*/
func (gf *File) Delete() error {
	c := GetClient()

	f, _, _, ferr := c.Repositories.GetContents(GetCtx(), gf.Owner, gf.Repo, gf.Keypath, nil)
	if ferr != nil {
		return ferr
	}
	if f == nil {
		// return errors.New("file is nil")
		// FILE not found
		return nil
	}

	message := "remove " + gf.Keypath
	_, _, err := c.Repositories.DeleteFile(GetCtx(), gf.Owner, gf.Repo, gf.Keypath, &github.RepositoryContentFileOptions{
		Message: &message,
		SHA:     f.SHA,
	})
	return err
}
