package nosh

import (
	"path"
	ut "utils"
)

/*
Env a
*/
type Env struct {
	FilePath string
}

/*
Create a
*/
func Create(filepath string) *Env {
	ut.Mkdirp(path.Dir(filepath))
	return &Env{filepath}
}

/*
Exists a
*/
func (env *Env) Exists() bool {
	return ut.IsFileExisted(env.FilePath)
}

/*
GetOrInstallNosh a
*/
func (env *Env) GetOrInstallNosh() string {
	if !env.Exists() {
		installNosh(env.FilePath)
	}
	return env.FilePath
}

/*
Upgrade a
*/
func (env *Env) Upgrade() string {
	installNosh(env.FilePath)
	return env.FilePath
}
