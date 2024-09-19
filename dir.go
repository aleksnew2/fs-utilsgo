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
	if strings.HasSuffix(path, "\\") {
		path = path + "\\"
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
// Found elements append to slice in this format:
//
// Directory format: dtest
//
// File format: ftest.txt
//
// If there's error, returns nil and error.
func ReadDir(path string) ([]string, error) {
	var slice []string
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			slice = append(slice, "d"+location)
		} else {
			slice = append(slice, "f"+location)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return slice, nil
}

// ReadDirQ reads directory and returns Dir object.
// Found elements append to slice in this format:
//
// Directory format: dtest
//
// File format: ftest.txt
//
// If there's error, returns nil and error.
func ReadDirQ(path string) (*Dir, error) {
	var children []string
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			children = append(children, "d"+location)
		} else {
			children = append(children, "f"+location)
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
			fmt.Printf("found dir: %v\n", location)
		} else {
			fmt.Printf("found file: %v\n", location)
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
// Found elements append to slice in this format:
//
// Directory format: dtest
//
// File format: ftest.txt
func ReadDirA(d *Dir) error {
	err := filepath.Walk(d.Path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			d.Children = append(d.Children, "d"+location)
		} else {
			d.Children = append(d.Children, "f"+location)
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// ReadDirD scans directory from specific path and outputs process.
// Generates random ID to identify an operation.
// Returns ID.
// If there's an error, then functions outputs error instead of panic.
func ReadDirD(path string) string {
	id := generateID(16)
	fmt.Printf("%v: starting scanning directory... (path: %v)\n", id, path)

	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			fmt.Printf("found directory: %v\n", location)
		} else {
			fmt.Printf("found file: %v\n", location)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("error while scanning: %v", err)
	}

	return id
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

// RemoveDirQ removes a directory from specific path.
func RemoveDirQ(path string) error {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	if !IsDirExists(path) {
		return fmt.Errorf("dir doesn't exists: (%v)", path)
	}

	if err := os.RemoveAll(path); err != nil {
		return err
	}

	return nil
}

// RemoveDirW removes a directory but from the structure Dir.
// If directory doesn't exist, then returns an error.
func RemoveDirW(d *Dir) error {
	if !IsDirExists(d.Path) {
		return fmt.Errorf("dir %v doesn't exist", d.Path)
	}

	if err := os.Remove(d.Path); err != nil {
		return err
	}

	emptyDirW(d)
	return nil
}

// RemoveDirA removes a directory but from the structure Dir.
// Returns directory's children.
// If directory doesn't exist, then returns an error.
func RemoveDirA(d *Dir) ([]string, error) {
	if !IsDirExists(d.Path) {
		return nil, fmt.Errorf("dir %v doesn't exist", d.Path)
	}

	if err := os.Remove(d.Path); err != nil {
		return nil, err
	}

	children := emptyDirQ(d)
	return children, nil
}

// Output outputs directory with fmt.Printf.
func (d Dir) Output() {
	fmt.Printf("Path: %v\nChildren: %v\n", d.Path, d.Children)
}

// MoveDir moves a directory from sourcePath to destinationPath.
// If the destination directory already exists, returns an error.
func MoveDir(sourcePath, destinationPath string) error {
	if IsDirExists(destinationPath) {
		return fmt.Errorf("directory %v already exists", destinationPath)
	}
	return os.Rename(sourcePath, destinationPath)
}

// ListFilesInDir lists all files in the specified directory.
// Returns a slice of file names and an error if any occurs.
func ListFilesInDir(path string) ([]string, error) {
	if !IsDirExists(path) {
		return nil, fmt.Errorf("directory %v does not exist", path)
	}

	var files []string
	err := filepath.Walk(path, func(location string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, location)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// RemoveEmptyDir removes an empty directory at the specified path.
// Returns an error if the directory is not empty or does not exist.
func RemoveEmptyDir(path string) error {
	if !IsDirExists(path) {
		return fmt.Errorf("directory %v does not exist", path)
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	if len(entries) > 0 {
		return fmt.Errorf("directory %v is not empty", path)
	}

	return os.Remove(path)
}
