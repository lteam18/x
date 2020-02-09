package ghkv

import (
	"encoding/json"
	"gh"
	"path"
	ut "utils"
	"vvkv"
)

/*
GitRes s
*/
type GitRes struct {
	Owner   string `json:"owner"`
	Repo    string `json:"repo"`
	ResPath string `json:"resPath"`
}

/*
CreateGitRes s
*/
func CreateGitRes(ghurl string) *GitRes {
	owner, repo, respath := ParseGHPrefix(ghurl)
	return &GitRes{
		Owner: owner, Repo: repo, ResPath: respath,
	}
}

var metaIdxPrefix = "meta.x-cmd.com/"

/*
GetPrivateKeyPath s
*/
func (res *GitRes) GetPrivateKeyPath() (string, string) {
	// return "APP_IDX/" + res.ResPath, "APP/" + res.ResPath
	return metaIdxPrefix + res.ResPath, res.ResPath
}

/*
GetPublicKeyPath s
*/
func (res *GitRes) GetPublicKeyPath() (string, string) {
	// return "docs/APP_IDX/" + res.ResPath, "docs/APP/" + res.ResPath
	return "docs/" + metaIdxPrefix + res.ResPath, "docs/" + res.ResPath
}

/*
GetPrivateGitFile s
*/
func (res *GitRes) GetPrivateGitFile() (*gh.File, *gh.File) {
	idx, app := res.GetPrivateKeyPath()

	idxGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: idx}
	appGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: app}
	return idxGF, appGF
}

/*
GetPublicGitFile s
*/
func (res *GitRes) GetPublicGitFile() (*gh.File, *gh.File) {
	idx, app := res.GetPublicKeyPath()

	idxGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: idx}
	appGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: app}
	return idxGF, appGF
}

/*
GetFilePathInCache s
*/
func (res *GitRes) GetFilePathInCache() (string, string) {
	sub := "/gh:" + res.Owner + ":" + res.Repo + "/" + res.ResPath
	idx := vvkv.AppIdxDirPath + sub
	app := vvkv.AppDirPath + sub
	return idx, app
}

/*
GetIdxAppPath s
*/
func (res *GitRes) GetIdxAppPath(isPublic bool) (string, string) {
	getter := res.GetPrivateKeyPath
	if isPublic {
		getter = res.GetPublicKeyPath
	}
	return getter()
}

// func (res *GitRes) getLocalAppIdxFilePath() string {
// 	return vvkv.AppIdxDirPath + "/gh:" + res.Owner + ":" + res.Repo + "/" + res.ResPath
// }

// func (res *GitRes) getLocalAppFilePath() string {
// 	return vvkv.AppDirPath + "/gh:" + res.Owner + ":" + res.Repo + "/" + res.ResPath
// }

/*
Meta s
*/
type Meta struct {
	CodeType string  `json:"codetype"`
	IsURL    bool    `json:"isURL"`
	URL      *string `json:"url"`
}

/*
MetaX s
*/
type MetaX struct {
	Meta       *Meta  `json:"meta"`
	IsPublic   bool   `json:"isPublic"`
	IdxKeyPath string `json:"idxKeyPath"`
	AppKeyPath string `json:"appKeyPath"`
}

// func (res *GitRes) GetMetaInCache() (*MetaX, error) {
// 	localAppIdxPath, _ := res.GetFilePathInCache()
// }

/*
GetMetaXInCache s
*/
func (res *GitRes) GetMetaXInCache(update bool) (*MetaX, error) {
	localAppIdxPath, _ := res.GetFilePathInCache()

	if update == false {
		if ut.IsFileExisted(localAppIdxPath) {
			log.Debug("Get " + localAppIdxPath)
			metaStr := ut.ReadFile(localAppIdxPath)
			var meta *MetaX
			if err := json.Unmarshal([]byte(metaStr), &meta); err != nil {
				return nil, err
			}
			return meta, nil
		}
	}

	log.Debug("Visiting Github to get meta")
	meta, err := res.GetMetaX()
	if err != nil {
		return nil, err
	}
	ut.Mkdirp(path.Dir(localAppIdxPath))
	ut.WriteFileSync(localAppIdxPath, ut.JS(meta))
	return meta, nil
}

/*
GetMetaX s
*/
func (res *GitRes) GetMetaX() (*MetaX, error) {
	log.Debug("Get Meta In Private")
	// TODO: If with authentiacation token, using api, instead of public url to access the resource.
	isPublic := false
	metaGF, _ := res.GetPrivateGitFile()
	metaStr, e1 := metaGF.Cat()

	if e1 != nil {
		isPublic = true
		metaGF, _ = res.GetPublicGitFile()
		metaStr, e1 = metaGF.Cat()

		if e1 != nil {
			return nil, e1
		}
	}

	var meta *Meta
	if e2 := json.Unmarshal([]byte(metaStr), &meta); e2 != nil {
		return nil, e2
	}

	appIdx, app := res.GetPrivateKeyPath()
	if isPublic {
		appIdx, app = res.GetPublicKeyPath()
	}

	return &MetaX{meta, isPublic, appIdx, app}, nil
}
