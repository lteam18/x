package ghkv

import "gh"

/*
InitGHRes i
*/
func InitGHRes(owner string) error {
	log.WithField("owner", owner).WithField("repo", CODE_REPO).Debug("Creating A Repo")
	err := gh.CreateRepo(false, owner, CODE_REPO)
	if err != nil {
		log.Error(err)
	} else {
		log.WithField("owner", owner).WithField("repo", CODE_REPO).Debug("Repo created")
	}

	// uploadBytesGHRes([]byte("echo hi"), "@gh/init", false, "bash")

	initGF := &gh.File{Owner: owner, Repo: CODE_REPO, Keypath: "docs/init"}
	initGF.Upload([]byte("echo hi"))
	return gh.EnablePages(owner, CODE_REPO)
}

func bindDomain(owner string) error {
	// gitowner.github.io/gitrepo => owner.x-cmd.com
	return nil
}
