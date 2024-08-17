package fs_utils

import (
	"io/fs"
	"os"
	"path/filepath"
)

// Dir is structure where contains information about
// specific directory.
// Should be initialized by functions ReadDir and ReadDirQ.
type Dir struct {
	Path     string
	Children []string
}

// ReadDir reads directory and returns string slice.
// If there's error, returns nil and error.
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

// ReadDirQ reads directory and returns Dir object.
// If there's error, returns nil and error.
func ReadDirQ(path string) (*Dir, error) {
	var children []string
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			children = append(children, "dir: "+location)
		} else {
			children = append(children, "file: "+location)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &Dir{path, children}, nil
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
