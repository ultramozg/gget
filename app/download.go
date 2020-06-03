package app

import (
	"io"
	"net/http"
	"os"

	"github.com/schollz/progressbar/v3"
)

func DownloadFile(url, file string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(file)
	if err != nil {
		return err
	}
	defer out.Close()

	bar := progressbar.DefaultBytes(
		resp.ContentLength,
		"downloading",
	)

	_, err = io.Copy(io.MultiWriter(out, bar), resp.Body)
	return err
}
