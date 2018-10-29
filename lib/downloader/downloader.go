package downloader

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mitchellh/ioprogress"
)

// GetPackage downloads a version of the game to a cache dir
func GetPackage(av AvailableVersion, cacheDir string) error {
	client := http.Client{Transport: http.DefaultTransport}

	response, err := client.Get(av.Target)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	progressR := &ioprogress.Reader{
		Reader: response.Body,
		Size:   response.ContentLength,
	}
	fileName := filepath.Join(cacheDir, av.File)
	fd, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	io.Copy(fd, progressR)
	return nil
}
