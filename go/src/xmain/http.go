package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	ut "utils"
)

/*
FileInfo f
*/
type FileInfo struct {
	Size int64  `json:"size"`
	Path string `json:"path"`
	Unix int64  `json:"ts"`
	Mode string `json:"mode"`
}

/*
FolderInfo f
*/
type FolderInfo struct {
	Info    FileInfo      `json:"info"`
	Folders []*FolderInfo `json:"folders"`
	Files   []FileInfo    `json:"files"`
}

func createFileInfo(path string, info os.FileInfo) FileInfo {
	return FileInfo{info.Size(), path, info.ModTime().Unix(), info.Mode().String()}
}

func createFolderInfo(path string, info os.FileInfo) *FolderInfo {
	return &FolderInfo{
		createFileInfo(path, info),
		make([]*FolderInfo, 0),
		make([]FileInfo, 0),
	}
}

/*
WalkDir s
*/
func WalkDir(root string) ([]FileInfo, error) {
	var files []FileInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, createFileInfo(path, info))
		}
		return nil
	})
	return files, err
}

/*
WalkDirInTree s
*/
func WalkDirInTree(root string) (*FolderInfo, error) {
	all := make([]*FolderInfo, 0)

	var last *FolderInfo
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if path == root {
			last = createFolderInfo(path, info)
			all = append(all, last)
			return nil
		}
		if !info.IsDir() {
			last.Files = append(last.Files, createFileInfo(path, info))
			return nil
		}

		folder := createFolderInfo(path, info)

		// filepath.Dir("dist") == "."
		for filepath.Dir(path) != last.Info.Path {
			// If the logic is right, it should be safe.
			all = all[0 : len(all)-1]
			last = all[len(all)-1]
		}

		last.Folders = append(last.Folders, folder)
		last = folder
		all = append(all, folder)

		return nil
	})
	return all[0], err
}

func createWalkTreeHandlerInFlatMode(dir string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		files, _ := WalkDir(dir)
		fmt.Fprintf(w, ut.JS(files))
	}
}

func createWalkTreeHandlerInTreeMode(dir string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		files, _ := WalkDirInTree(dir)
		fmt.Fprintf(w, ut.JS(files))
	}
}

/*
Serve s
*/
func Serve(port int, dir string) {
	// dir, _ = filepath.Abs(dir)

	handler := http.FileServer(http.Dir(dir))
	http.Handle("/", handler) // TODO: When nav ready, use nav
	http.Handle("/raw/", http.StripPrefix("/raw/", handler))
	// http.Handle("/nav", http.FileServer(http.Dir(dir)))
	// http.Handle("/tree", http.FileServer(http.Dir(dir)))
	http.HandleFunc("/flat-json", createWalkTreeHandlerInFlatMode(dir))
	http.HandleFunc("/tree-json", createWalkTreeHandlerInTreeMode(dir))

	a := `Listening at port %d 
Serving folder: %s
  => http://localhost:%d will bring you a navigator view, just as visitting http://localhost:%d/nav
  => http://localhost:%d/raw will bring you a plain navigator
  => http://localhost:%d/flat-json will return a list of all files in json format
  => http://localhost:%d/tree-json will return a json string where you files is organizing in a tree 
`

	fmt.Printf(a, port, dir, port, port, port, port, port)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}
