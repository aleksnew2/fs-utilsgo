package fs_utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// ReadDir reads directory and returns string slice.
func ReadDir(path string) ([]string, error) {
	var slice []string
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			slice = append(slice, "dir: "+location)
		} else {
			slice = append(slice, "file"+location)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return slice, nil
}

// CreateDir creates directory to specific path with os.Mkdir.
func CreateDir(path string) error {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// CreateDirQ creates directory to specific path with os.MkdirAll.
func CreateDirQ(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
