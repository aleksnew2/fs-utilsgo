package fs_utils

import (
	"bufio"
	"fmt"
	"os"
)

// File is a structure with information about file.
// File should be initialized by functions
// CreateFileQ or CreateFileW.
// Can be removed by RemoveFileW or RemoveFileA.
type File struct {
	Path    string
	Content FileLines
}

// FileLines contains lines of specific file.
// Should be initialized by GetFileContent.
type FileLines []string

// emptyFileW makes property f empty.
func emptyFileW(f *File) {
	f.Path = ""
	f.Content = nil
}

// emptyFileQ makes property f empty
// and returns content of file.
func emptyFileQ(f *File) []string {
	lastContent := f.Content
	f.Path = ""
	f.Content = nil
	return lastContent
}

// IsFileExists checks file existence.
func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// GetFile returns path file. Checks their availability.
// If file doesn't exist, return empty string and error.
func GetFile(path string) (string, error) {
	if !IsFileExists(path) {
		return "", fmt.Errorf("%v doesn't exist", path)
	}
	return path, nil
}

// CreateFileQ creates a file at a specific path.
// If the file already exists, then returns an error.
func CreateFileQ(path string) (*File, error) {
	if IsFileExists(path) {
		return nil, fmt.Errorf("file %v already exists", path)
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return &File{path, []string{""}}, nil
}

// CreateFileW creates a file at a specific path,
// then writes content to the file.
// Every element of content is a new line.
// If the file already exists, then returns an error.
func CreateFileW(path string, content ...string) (*File, error) {
	if IsFileExists(path) {
		return nil, fmt.Errorf("file %v already exists", path)
	}

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

// RemoveFileQ removes a file at a specific path.
// If it couldn't find the file, then returns an error.
func RemoveFileQ(path string) error {
	if !IsFileExists(path) {
		return fmt.Errorf("file %v does not exist", path)
	}

	if err := os.Remove(path); err != nil {
		return err
	}

	return nil
}

// RemoveFileW removes a file from the structure File.
// If it couldn't find the file, then returns an error.
func RemoveFileW(f *File) error {
	if !IsFileExists(f.Path) {
		return fmt.Errorf("file %v does not exist", f.Path)
	}

	if err := os.Remove(f.Path); err != nil {
		return err
	}

	emptyFileW(f)
	return nil
}

// RemoveFileA removes a file from the structure File.
// Returns the content from the file.
// If it couldn't find the file, then returns an empty string slice and an error.
func RemoveFileA(f *File) ([]string, error) {
	if !IsFileExists(f.Path) {
		return nil, fmt.Errorf("file %v does not exist", f.Path)
	}

	if err := os.Remove(f.Path); err != nil {
		return nil, err
	}

	content := emptyFileQ(f)
	return content, nil
}

// OutputFileContent outputs file content by the line.
// Example:
//
// 1. HI!
//
// 2. BYE!
//
// If there's error, function panics.
func OutputFileContent(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)

	for i := 0; i < len(scanner.Text()); i++ {
		for scanner.Scan() {
			i++
			fmt.Printf("%v. %v", i, scanner.Text())
		}
	}
}

// GetFileContent returns slice of content from specific file.
// Every element of slice marked as one line.
// If there's an error, function returns nil and error.
func GetFileContent(path string) (FileLines, error) {
	var lines FileLines

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

// WriteContent writes content to file.
// If it couldn't, returns error.\
func WriteContent(path string, content ...string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := bufio.NewWriter(file)

	defer func(writer *bufio.Writer) {
		_ = writer.Flush()
	}(writer)

	for _, line := range content {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

// Output outputs lines.
// If there aren't lines, outputs error.
func (fl FileLines) Output() {
	if len(fl) == 0 {
		fmt.Printf("FileLines.Output: there aren't lines")
	}

	for i := 0; i < len(fl); {
		for _, line := range fl {
			i++
			fmt.Printf("%v. %v", i, line)
		}
	}
}
