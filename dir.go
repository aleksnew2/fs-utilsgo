package fs_utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// Dir is structure where contains information about
// specific directory.
// Should be initialized by functions ReadDir and ReadDirQ,
// or CreateDirW.
type Dir struct {
	Path     string
	Children []string
}

// emptyDirW makes property d empty.
func emptyDirW(d *Dir) {
	d.Path = ""
	d.Children = nil
}

// emptyDirQ makes property d empty
// and returns directory's children.
func emptyDirQ(d *Dir) []string {
	lastChildren := d.Children
	d.Path = ""
	d.Children = nil
	return lastChildren
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

// ReadDirW reads directory and outputs content with fmt.Printf.
func ReadDirW(path string) error {
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Printf("dir: %v\n", location)
		} else {
			fmt.Printf("file: %v\n", location)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// ReadDirA works same as ReadDir etc.
// But it reads directory and puts children
// to d.Children.
func ReadDirA(d *Dir) error {
	err := filepath.Walk(d.Path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			d.Children = append(d.Children, "dir: "+location)
		} else {
			d.Children = append(d.Children, "file: "+location)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
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

// CreateDirW creates directory to specific path with os.Mkdir.
// Then returns Dir object.
func CreateDirW(path string) (*Dir, error) {
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &Dir{Path: path}, nil
}

// Output outputs directory with fmt.Printf.
func (d Dir) Output() {
	fmt.Printf("Path: %v\nChildren: %v\n", d.Path, d.Children)
}
