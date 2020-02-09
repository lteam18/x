package gh

import (
	"github.com/google/go-github/github"
)

/*
CreateRepo c
*/
func CreateRepo(public bool, owner, repo string) error {
	_, _, err := GetClient().Repositories.Get(GetCtx(), owner, repo)
	autoInit := true
	isPrivate := !public
	if err != nil {
		_, _, err := GetClient().Repositories.Create(GetCtx(), "", &github.Repository{
			Name:     &repo,
			AutoInit: &autoInit,
			Private:  &isPrivate,
		})

		return err
	}

	_, _, e1 := GetClient().Repositories.Edit(GetCtx(), owner, repo, &github.Repository{
		Name:     &repo,
		AutoInit: &autoInit,
		Private:  &isPrivate,
	})

	return e1
}

/*
EnablePages e
*/
func EnablePages(owner, repo string) error {
	GetClient().Repositories.DisablePages(GetCtx(), owner, repo)
	return getGitClient2().EnablePage(owner, repo, "master", "/docs")
}

/*
DeleteRepo e
*/
func DeleteRepo(owner, repo string) error {
	_, err := GetClient().Repositories.Delete(GetCtx(), owner, repo)
	return err
}
