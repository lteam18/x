package ghkv

import (
	"fmt"
	"gh"
	"strings"
)

type item struct {
	name string
	size int
}

/*
Item s
*/
type Item struct {
	Name     string `json:"name"`
	Size     int    `json:"size"`
	IsPublic bool   `json:"isPublic"`
}

/*
LsQueryStr l
*/
func LsQueryStr(query string) *string {
	qs := lsQuery(query)
	if qs == nil {
		return nil
	}
	s := []string{}
	for _, q := range *qs {
		mode := "private"
		if q.IsPublic {
			mode = "public"
		}
		s = append(s, fmt.Sprintf("%v\t%v\t%v", q.Name, q.Size, mode))
	}
	ret := strings.Join(s, "\n")
	return &ret
}

var publicIdxLen = len("docs/" + metaIdxPrefix)

func lsQuery(originalQuery string) *[]Item {
	owner, repo, query := ParseGHPrefix(originalQuery)

	log.WithFields(map[string]interface{}{
		"query":         query,
		"owner":         owner,
		"repo":          repo,
		"originalQuery": originalQuery,
	}).Debug("lsQuery()")

	all := ls(owner, repo, "")

	ret := make(map[string]Item)

	if all == nil {
		return &[]Item{}
	}

	for _, s := range *all {
		if strings.HasPrefix(s.name, "docs/") {
			if strings.HasPrefix(s.name, "docs/"+metaIdxPrefix) {
				name := s.name[publicIdxLen:]
				if !strings.HasPrefix(name, query) {
					continue
				}
				if _, ok := ret[name]; ok == false {
					ret[name] = Item{
						name,
						s.size,
						true,
					}
				}
			} else {
				name := s.name[5:]
				if !strings.HasPrefix(name, query) {
					continue
				}
				if it, ok := ret[name]; ok {
					it.Size = s.size
				} else {
					ret[name] = Item{
						name,
						s.size,
						true,
					}
				}
			}
		} else {
			if strings.HasPrefix(s.name, metaIdxPrefix) {
				name := s.name[len(metaIdxPrefix):]
				if !strings.HasPrefix(name, query) {
					continue
				}
				if _, ok := ret[name]; ok == false {
					ret[name] = Item{
						name,
						s.size,
						false,
					}
				}
			} else {
				name := s.name
				if !strings.HasPrefix(name, query) {
					continue
				}
				if it, ok := ret[name]; ok {
					it.Size = s.size
				} else {
					ret[name] = Item{
						name,
						s.size,
						false,
					}
				}
			}
		}
	}

	res := []Item{}
	for _, v := range ret {
		res = append(res, v)
	}

	return &res
}

// TODO: using loop to iterate all files.
func ls(owner, repo, query string) *[]item {
	log.WithFields(map[string]interface{}{
		"query": query,
		"owner": owner,
		"repo":  repo,
	}).Debug("ls()")

	_, dirs, _, err := gh.GetClient().Repositories.GetContents(gh.GetCtx(), owner, repo, query, nil)
	if err != nil {
		return nil
	}

	ret := []item{}
	for _, file := range dirs {
		if file.GetType() == "dir" {
			res := ls(owner, repo, file.GetPath())
			if res != nil {
				ret = append(ret, *res...)
			}
			continue
		}
		ret = append(ret, item{
			name: file.GetPath(),
			size: file.GetSize(),
		})
		// fmt.Sprintf("%s\t%d", *file.Path, *file.Size))
	}

	return &ret
}
