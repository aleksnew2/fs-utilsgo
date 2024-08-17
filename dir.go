package fs_utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Dir is structure where contains information about
// specific directory.
// Should be initialized by functions ReadDirQ,
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

func IsDirExists(path string) bool {
	if strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// GetDir returns directory from specific path.
// If directory doesn't exist, then returns empty string and error.
func GetDir(path string) (string, error) {
	if !IsDirExists(path) {
		return "", fmt.Errorf("dir doesn't exist (%v)", path)
	}

	return path, nil
}

// GetDirQ returns directory from specific path.
// It takes path from property d, who inherited by structure Dir.
// Then, reads directory and put children to d.Children.
// If directory doesn't exist, then returns empty d object and error.
func GetDirQ(d *Dir) (*Dir, error) {
	if !IsDirExists(d.Path) {
		return nil, fmt.Errorf("dir doesn't exist (%v)", d.Path)
	}

	fsElements, err := ReadDir(d.Path)
	if err != nil {
		return nil, err
	}

	copy(d.Children, fsElements)
	return &Dir{Path: d.Path, Children: fsElements}, nil
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

// RemoveDir removes directory from specific path.
func RemoveDir(path string) error {
	if !strings.HasSuffix(path, "/") {

	}

	if !IsDirExists(path) {
		return fmt.Errorf("dir doesn't exists: (%v)", path)
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

// Output outputs directory with fmt.Printf.
func (d Dir) Output() {
	fmt.Printf("Path: %v\nChildren: %v\n", d.Path, d.Children)
}
