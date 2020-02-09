package ghkv

import (
	"encoding/json"
	"gh"
	"os"
	ut "utils"
)

/*
UploadURL s
*/
func (res *GitRes) UploadURL(srcURL string, isPublic bool, codetype string) error {
	creator := res.GetPrivateGitFile
	if isPublic {
		creator = res.GetPublicGitFile
	}
	idxGF, _ := creator()

	meta := &Meta{codetype, true, &srcURL}
	st, _ := json.Marshal(meta)

	return idxGF.Upload(st)
}

/*
GetCodeTypeAndTempURL s
*/
func (res *GitRes) GetCodeTypeAndTempURL() (*string, *string, error) {
	meta, e1 := res.GetMetaX()

	if e1 != nil {
		return nil, nil, e1
	}

	if !meta.IsPublic {
		// f := res.GetPrivateGitFile
		// if meta.IsPublic {
		// 	f = res.GetPublicGitFile
		// }

		_, app := res.GetPrivateGitFile()
		url, err := app.GetTempURL()
		if err != nil {
			return nil, nil, err
		}
		return &meta.Meta.CodeType, &url, nil
	}

	_, app := res.GetPrivateGitFile()
	url := app.GetGHPageURL()

	return &meta.Meta.CodeType, &url, nil

}

/*
UploadBytes s
*/
func (res *GitRes) UploadBytes(bytes []byte, isPublic bool, codetype string) error {
	idx, app := res.GetIdxAppPath(isPublic)
	// res.DeleteFiles(idx, app)

	// path not startswith /
	metaStr, _ := json.Marshal(&Meta{codetype, false, nil})

	idxGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: idx}
	if err := idxGF.Upload(metaStr); err != nil {
		log.Debug(err)
	}

	appGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: app}
	if err := appGF.Upload(bytes); err != nil {
		log.Debug(err)
	}
	return nil
}

/*
UploadFile s
*/
func (res *GitRes) UploadFile(filepath string, isPublic bool, codetype string) error {
	return res.UploadBytes(ut.ReadFileToBytes(filepath), isPublic, codetype)
}

// TODO: a lot of things to do to make sure it works

/*
SetAccess s
*/
func (res *GitRes) SetAccess(dstIsPublic bool) error {
	meta, err := res.GetMetaX()

	if err != nil {
		return err
	}

	if meta.IsPublic == dstIsPublic {
		return nil
	}

	srcIdxGF, srcAppGF := res.GetPublicGitFile()
	dstIdx, dstApp := res.GetPrivateKeyPath()

	if meta.IsPublic == false {
		srcIdxGF, srcAppGF = res.GetPrivateGitFile()
		dstIdx, dstApp = res.GetPublicKeyPath()
	}

	if err := srcAppGF.CopyTo(dstApp); err != nil {
		return err
	}

	if err := srcIdxGF.CopyTo(dstIdx); err != nil {
		return err
	}

	if derr2 := srcAppGF.Delete(); derr2 != nil {
		log.Error(derr2)
	}

	if derr1 := srcIdxGF.Delete(); derr1 != nil {
		log.Error(derr1)
	}

	idxLP, _ := res.GetFilePathInCache()
	os.Remove(idxLP)
	// os.Remove(appLP)

	return nil
}

/*
DeleteFiles s
*/
func (res *GitRes) DeleteFiles(idx string, app string) error {

	idxGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: idx}
	if err := idxGF.Delete(); err != nil {
		log.Debug(err)
	}

	appGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: app}
	if err := appGF.Delete(); err != nil {
		log.Debug(err)
	}

	idxLP, appLP := res.GetFilePathInCache()
	os.Remove(idxLP)
	os.Remove(appLP)

	return nil
}

/*
Delete s
*/
func (res *GitRes) Delete() error {

	log.WithField("gitres", res).WithField("meta", *res).Debug("Delete")
	meta, err := res.GetMetaX()

	if err != nil {
		return err
	}

	idxGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: meta.IdxKeyPath}
	err1 := idxGF.Delete()
	if err1 != nil {
		return err1
	}

	appGF := &gh.File{Owner: res.Owner, Repo: res.Repo, Keypath: meta.AppKeyPath}
	err2 := appGF.Delete()
	if err2 != nil {
		return err2
	}

	idxLP, appLP := res.GetFilePathInCache()
	os.Remove(idxLP)
	os.Remove(appLP)

	return nil
}
