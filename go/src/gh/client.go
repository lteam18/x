package gh

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

/*
GetCtx g
*/
func GetCtx() context.Context {
	ctx := context.Background()
	return ctx
}

/*
GetClient g
*/
func GetClient() *github.Client {

	// token := "9f3109a6b792f8ca65d3982d8753e9a7e34da700"
	token := GetToken()

	if token == nil {
		return github.NewClient(oauth2.NewClient(GetCtx(), nil))
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token.Token},
	)
	tc := oauth2.NewClient(GetCtx(), ts)
	client := github.NewClient(tc)
	return client
}
