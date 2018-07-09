package util

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type PathFunc func(f *zip.File) string

func GetURL(url string, attempts int) (*http.Response, int, error) {
	var err error
	var resp *http.Response
	var size int
	for i := 0; i < attempts; i++ {
		resp, err = http.Get(url)
		if err == nil {
			size = int(resp.ContentLength)
			if size > 0 {
				return resp, size, nil
			}
			resp.Body.Close()
		}
		time.Sleep(time.Second)
	}
	return nil, 0, err
}

func Download(url, dir, name string) (string, error) {
	r, _, e := GetURL(url, 5)
	if e != nil {
		return "", e
	}
	defer r.Body.Close()

	e = os.MkdirAll(dir, os.ModePerm)
	if e != nil {
		return "", e
	}

	p := filepath.Join(dir, name)
	f, e := os.Create(p)
	if e != nil {
		return "", e
	}
	defer f.Close()

	_, e = io.Copy(f, r.Body)
	if e != nil {
		return "", e
	}

	return p, nil
}

func Extract(path, dir string, fn PathFunc) error {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer zr.Close()

	for _, f := range zr.File {
		if f.FileInfo().IsDir() {
			continue
		}

		path := fn(f)
		if path == "" {
			continue
		}

		fmt.Print(path)
		path = filepath.Join(dir, path)
		fmt.Println(" -> ", path)

		if Exists(path) {
			// file exists
			continue
		}

		e := os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if e != nil {
			fmt.Println("mkdir", e)
			continue
		}

		to, e := os.Create(path)
		if e != nil {
			fmt.Println("create", e)
			// todo
			continue
		}

		from, e := f.Open()
		if e != nil {
			fmt.Println("open", e)
			// todo
			continue
		}

		_, e = io.Copy(to, from)
		if e != nil {
			fmt.Println("copy", e)
			// todo
		}
	}

	return nil
}

func Move(from, to string) error {
	e := os.MkdirAll(filepath.Dir(to), os.ModePerm)
	if e != nil {
		return e
	}
	return os.Rename(from, to)
}

func Exists(path string) bool {
	_, e := os.Stat(path)
	return e == nil
}
