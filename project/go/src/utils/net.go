package ut

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/*
Download d
TODO: Provide progressbar and become stoppable
*/
func Download(url string, dstPath string) error {
	resp, _ := http.Get(url)

	temppath, temppathErr := ioutil.TempFile(os.TempDir(), "tmp*")
	if temppathErr != nil {
		return temppathErr
	}

	if _, err := io.Copy(temppath, resp.Body); err != nil {
		return err
	}

	if err := CopyFile(temppath.Name(), dstPath); err != nil {
		return err
	}

	if err := os.Remove(temppath.Name()); err != nil {
		return err
	}

	return nil
}
