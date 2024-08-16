package fs_utils

import (
	"os"
)

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
