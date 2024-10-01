package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileSystem interface {
	Open(name string) (io.ReadCloser, error)
}

type RealFileSystem struct{}

func (RealFileSystem) Open(name string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(".", name))
}

func readFile(fs FileSystem, name string) string {
	f, err := fs.Open(filepath.Join(".", name))
	var buf strings.Builder
	_, err = io.Copy(&buf, f)
	if err != nil {
		log.Fatal(err)
	}
	return buf.String()
}
