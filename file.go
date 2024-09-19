package fs_utils

import (
	"bufio"
	"fmt"
	"io"
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
func emptyFileQ(f *File) FileLines {
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
func CreateFileW(path string, content FileLines) (*File, error) {
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

// CreateFileA creates a file at a specific path,
// then writes content to the file.
// Every element of content is a new line.
// If the file already exists, then returns an error.
func CreateFileA(path string, content FileLines) error {
	if IsFileExists(path) {
		return fmt.Errorf("file %v already exists", path)
	}

	file, err := os.Create(path)
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

// CreateFileR creates a file at a specific path.
// If the file already exists, then returns an error.
func CreateFileR(path string) error {
	if IsFileExists(path) {
		return fmt.Errorf("file %v already exists", path)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	return nil
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
func RemoveFileA(f *File) (FileLines, error) {
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

	i := 1
	for scanner.Scan() {
		fmt.Printf("%v. %v\n", i, scanner.Text())
		i++
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
// If it couldn't, returns error.
func WriteContent(path string, content FileLines) error {
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
		fmt.Printf("FileLines.Output: there aren't lines\n")
		return
	}

	for i := 0; i < len(fl); i++ {
		fmt.Printf("%v. %v\n", i+1, fl[i])
	}
}

// RenameFile renames a file from oldPath to newPath.
// If the newPath already exists, returns an error.
func RenameFile(oldPath, newPath string) error {
	if IsFileExists(newPath) {
		return fmt.Errorf("file %v already exists", newPath)
	}
	return os.Rename(oldPath, newPath)
}

// CopyFile copies a file from source to destination.
// If the destination file already exists, returns an error.
func CopyFile(source, destination string) error {
	if IsFileExists(destination) {
		return fmt.Errorf("file %v already exists", destination)
	}

	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()

	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer output.Close()

	_, err = io.Copy(output, input)
	return err
}

// AppendToFile appends content to an existing file.
// If the file doesn't exist, returns an error.
func AppendToFile(path string, content FileLines) error {
	if !IsFileExists(path) {
		return fmt.Errorf("file %v does not exist", path)
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range content {
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
