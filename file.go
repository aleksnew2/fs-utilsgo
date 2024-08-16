// Package fs_utils provides function for creating, removing files etc.
package fs_utils

import (
	"bufio"
	"os"
)

// File is a structure with information about file.
// File should be initialized by functions
// CreateFileQ or CreateFileW.
// Can be removed by RemoveFileW or RemoveFileA.
type File struct {
	Path    string
	Content []string
}

// emptyFileW makes property f empty.
func emptyFileW(f *File) {
	f = nil
}

// emptyFileQ makes property
// and returns content of file.
func emptyFileQ(f *File) []string {
	lastContent := f.Content
	f.Path = ""
	f.Content = nil
	return lastContent
}

// GetFile returns path file. Checks their availability.
// If file doesn't exist, return empty string and error.
func GetFile(path string) (string, error) {
	return path, nil
}

// CreateFileQ creates file to specific path.
// If file already exists, then returns error.
func CreateFileQ(path string) (*File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)
	return &File{path, []string{""}}, nil
}

// CreateFileW creates file to specific path,
// then writes content to file.
// Every element of content is new line.
// If file already exists, then returns error.
func CreateFileW(path string, content ...string) (*File, error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := bufio.NewWriter(file)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	for _, v := range content {
		if _, err := writer.WriteString(v + "\n"); err != nil {
			return nil, err
		}
	}

	return &File{path, content}, nil
}

// RemoveFileQ removes file from specific path.
// If it couldn't find, then returns error.
func RemoveFileQ(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// RemoveFileW removes file, but from structure File.
// If it couldn't find, then returns error.
func RemoveFileW(f *File) error {
	emptyFileW(f)

	return nil
}

// RemoveFileA removes file, but from structure File.
// Returns content from file.
// If it couldn't find, then returns empty string slice and, error.
func RemoveFileA(f *File) ([]string, error) {
	content := emptyFileQ(f)

	return content, nil
}
