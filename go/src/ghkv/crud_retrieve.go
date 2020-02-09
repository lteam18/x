package ghkv

import "gh"

func (res *GitRes) retrievePublic(fileCachePath string, update bool) error {
	if gh.GetToken() != nil {
		_, gf := res.GetPublicGitFile()
		if err := gf.DownloadWithLocalCache(fileCachePath, update); err == nil {
			return nil
		}
		// Sometimes, error occur when access key not found
	}

	_, gf := res.GetPrivateGitFile()
	if err := gf.DownloadFromGithubPageWithCache(fileCachePath, update); err != nil {
		return err
	}

	return nil
}

func (res *GitRes) retrievePrivate(fileCachePath string, update bool) error {
	_, gf := res.GetPrivateGitFile()
	if err := gf.DownloadWithLocalCache(fileCachePath, update); err != nil {
		return err
	}
	return nil
}

/*
Retrieve s
*/
func (res *GitRes) Retrieve(update bool) (*Meta, *string, error) {
	localAppIdxPath, fileCachePath := res.GetFilePathInCache()

	log.WithFields(map[string]interface{}{
		"gitres":        res,
		"localFilePath": localAppIdxPath,
		"update":        update,
	}).Debugf("getFileWithLocalCache()")

	meta, metaError := res.GetMetaXInCache(update)

	if metaError != nil {
		// Under the situation where there is no index file
		// using empty string. It is a magic string. But have no better option.
		emptyMeta := &Meta{CodeType: "", IsURL: false, URL: nil}

		if err := res.retrievePublic(fileCachePath, update); err == nil {
			return emptyMeta, &fileCachePath, nil
		}

		if err := res.retrievePrivate(fileCachePath, update); err == nil {
			return emptyMeta, &fileCachePath, nil
		}

		return nil, nil, metaError
	}

	if meta.Meta.URL != nil {
		return meta.Meta, &localAppIdxPath, nil
	}

	if meta.IsPublic {
		if err := res.retrievePublic(fileCachePath, update); err != nil {
			return nil, nil, err
		}
		return meta.Meta, &fileCachePath, nil
	}

	if err := res.retrievePrivate(fileCachePath, update); err != nil {
		return nil, nil, err
	}
	return meta.Meta, &fileCachePath, nil
}
