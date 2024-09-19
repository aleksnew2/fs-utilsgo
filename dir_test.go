package fs_utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestIsDirExists(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Test existing directory
	if !IsDirExists(tempDir) {
		t.Errorf("expected directory to exist: %v", tempDir)
	}

	// Test non-existing directory
	nonExistentDir := filepath.Join(tempDir, "nonexistent")
	if IsDirExists(nonExistentDir) {
		t.Errorf("expected directory to not exist: %v", nonExistentDir)
	}
}

func TestGetDir(t *testing.T) {
	tempDir := t.TempDir()

	// Test existing directory
	dir, err := GetDir(tempDir)
	if err != nil || dir != tempDir {
		t.Errorf("expected to get directory: %v, got: %v, error: %v", tempDir, dir, err)
	}

	// Test non-existing directory
	nonExistentDir := filepath.Join(tempDir, "nonexistent")
	dir, err = GetDir(nonExistentDir)
	if err == nil || dir != "" {
		t.Errorf("expected error for non-existing directory, got: %v, error: %v", dir, err)
	}
}

func TestGetDirQ(t *testing.T) {
	tempDir := t.TempDir()
	dir := &Dir{Path: tempDir}

	// Test existing directory
	d, err := GetDirQ(dir)
	if err != nil {
		t.Errorf("expected to get directory: %v, error: %v", tempDir, err)
	}
	if d.Path != tempDir {
		t.Errorf("expected directory path to be: %v, got: %v", tempDir, d.Path)
	}

	// Test non-existing directory
	nonExistentDir := filepath.Join(tempDir, "nonexistent")
	dir.Path = nonExistentDir
	d, err = GetDirQ(dir)
	if err == nil {
		t.Errorf("expected error for non-existing directory, got: %v", d)
	}
}

func TestReadDir(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test reading directory
	elements, err := ReadDir(tempDir)
	if err != nil {
		t.Errorf("expected to read directory: %v, error: %v", tempDir, err)
	}

	expectedElements := []string{
		"f" + file1,
		"f" + file2,
		"d" + subdir,
	}
	for _, ee := range expectedElements {
		found := false
		for _, e := range elements {
			if e == ee {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected element: %v to be found", ee)
		}
	}
}

func TestReadDirQ(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test reading directory
	d, err := ReadDirQ(tempDir)
	if err != nil {
		t.Errorf("expected to read directory: %v, error: %v", tempDir, err)
	}

	expectedElements := []string{
		"f" + file1,
		"f" + file2,
		"d" + subdir,
	}
	for _, ee := range expectedElements {
		found := false
		for _, e := range d.Children {
			if e == ee {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected element: %v to be found", ee)
		}
	}
}

func TestReadDirW(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test reading directory
	err := ReadDirW(tempDir)
	if err != nil {
		t.Errorf("expected to read directory: %v, error: %v", tempDir, err)
	}
}

func TestReadDirA(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test reading directory
	d := &Dir{Path: tempDir}
	err := ReadDirA(d)
	if err != nil {
		t.Errorf("expected to read directory: %v, error: %v", tempDir, err)
	}

	expectedElements := []string{
		"f" + file1,
		"f" + file2,
		"d" + subdir,
	}
	for _, ee := range expectedElements {
		found := false
		for _, e := range d.Children {
			if e == ee {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected element: %v to be found", ee)
		}
	}
}

func TestReadDirD(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test reading directory
	id := ReadDirD(tempDir)
	if id == "" {
		t.Errorf("expected to get a valid ID, got: %v", id)
	}
}

func TestCreateDir(t *testing.T) {
	tempDir := t.TempDir()
	newDir := filepath.Join(tempDir, "newdir")

	// Test creating a new directory
	err := CreateDir(newDir)
	if err != nil {
		t.Errorf("expected to create directory: %v, error: %v", newDir, err)
	}

	// Verify directory was created
	if !IsDirExists(newDir) {
		t.Errorf("expected directory to exist: %v", newDir)
	}
}

func TestCreateDirQ(t *testing.T) {
	tempDir := t.TempDir()
	newDir := filepath.Join(tempDir, "newdir", "subdir")

	// Test creating a new directory
	err := CreateDirQ(newDir)
	if err != nil {
		t.Errorf("expected to create directory: %v, error: %v", newDir, err)
	}

	// Verify directory was created
	if !IsDirExists(newDir) {
		t.Errorf("expected directory to exist: %v", newDir)
	}
}

func TestCreateDirW(t *testing.T) {
	tempDir := t.TempDir()
	newDir := filepath.Join(tempDir, "newdir")

	// Test creating a new directory
	d, err := CreateDirW(newDir)
	if err != nil {
		t.Errorf("expected to create directory: %v, error: %v", newDir, err)
	}

	// Verify directory was created
	if !IsDirExists(newDir) {
		t.Errorf("expected directory to exist: %v", newDir)
	}

	// Verify directory path is correct
	if d.Path != newDir {
		t.Errorf("expected directory path to be: %v, got: %v", newDir, d.Path)
	}
}

func TestRemoveDirQ(t *testing.T) {
	tempDir := t.TempDir()

	// Test removing an existing directory
	err := RemoveDirQ(tempDir)
	if err != nil {
		t.Errorf("expected to remove directory: %v, error: %v", tempDir, err)
	}

	// Verify directory was removed
	if IsDirExists(tempDir) {
		t.Errorf("expected directory to not exist: %v", tempDir)
	}
}

func TestRemoveDirW(t *testing.T) {
	tempDir := t.TempDir()
	dir := &Dir{Path: tempDir}

	// Test removing an existing directory
	err := RemoveDirW(dir)
	if err != nil {
		t.Errorf("expected to remove directory: %v, error: %v", tempDir, err)
	}

	// Verify directory was removed
	if IsDirExists(tempDir) {
		t.Errorf("expected directory to not exist: %v", tempDir)
	}

	// Verify directory is empty
	if dir.Path != "" || len(dir.Children) != 0 {
		t.Errorf("expected directory to be empty, got: %v", dir)
	}
}

func TestRemoveDirA(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test removing an existing directory
	dir := &Dir{Path: tempDir}
	children, err := RemoveDirA(dir)
	if err != nil {
		t.Errorf("expected to remove directory: %v, error: %v", tempDir, err)
	}

	// Verify directory was removed
	if IsDirExists(tempDir) {
		t.Errorf("expected directory to not exist: %v", tempDir)
	}

	// Verify children are returned
	expectedChildren := []string{
		"f" + file1,
		"f" + file2,
		"d" + subdir,
	}
	for _, ec := range expectedChildren {
		found := false
		for _, c := range children {
			if c == ec {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected child: %v to be returned", ec)
		}
	}
}

func TestDirOutput(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")
	subdir := filepath.Join(tempDir, "subdir")

	// Create test files and subdirectory
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)
	os.Mkdir(subdir, os.ModePerm)

	// Test outputting directory
	d := &Dir{Path: tempDir}
	err := ReadDirA(d)
	if err != nil {
		t.Errorf("expected to read directory: %v, error: %v", tempDir, err)
	}

	// Capture output
	var output []string
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
		w.Close()
		r.Close()
	}()

	d.Output()

	// Read output
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		output = append(output, scanner.Text())
	}

	// Verify output
	expectedOutput := []string{
		fmt.Sprintf("Path: %v", tempDir),
		fmt.Sprintf("Children: %v", d.Children),
	}
	for i, eo := range expectedOutput {
		if output[i] != eo {
			t.Errorf("expected output: %v, got: %v", eo, output[i])
		}
	}
}

func TestMoveDir(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "source")
	destinationDir := filepath.Join(tempDir, "destination")

	// Create source directory
	os.Mkdir(sourceDir, os.ModePerm)

	// Test moving directory
	err := MoveDir(sourceDir, destinationDir)
	if err != nil {
		t.Errorf("expected to move directory: %v, error: %v", sourceDir, err)
	}

	// Verify source directory was removed
	if IsDirExists(sourceDir) {
		t.Errorf("expected source directory to not exist: %v", sourceDir)
	}

	// Verify destination directory was created
	if !IsDirExists(destinationDir) {
		t.Errorf("expected destination directory to exist: %v", destinationDir)
	}
}

func TestListFilesInDir(t *testing.T) {
	tempDir := t.TempDir()
	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(tempDir, "file2.txt")

	// Create test files
	os.WriteFile(file1, []byte("test"), 0644)
	os.WriteFile(file2, []byte("test"), 0644)

	// Test listing files in directory
	files, err := ListFilesInDir(tempDir)
	if err != nil {
		t.Errorf("expected to list files in directory: %v, error: %v", tempDir, err)
	}

	expectedFiles := []string{file1, file2}
	for _, ef := range expectedFiles {
		found := false
		for _, f := range files {
			if f == ef {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected file: %v to be listed", ef)
		}
	}
}

func TestRemoveEmptyDir(t *testing.T) {
	tempDir := t.TempDir()
	emptyDir := filepath.Join(tempDir, "emptydir")

	// Create an empty directory
	os.Mkdir(emptyDir, os.ModePerm)

	// Test removing an empty directory
	err := RemoveEmptyDir(emptyDir)
	if err != nil {
		t.Errorf("expected to remove empty directory: %v, error: %v", emptyDir, err)
	}

	// Verify directory was removed
	if IsDirExists(emptyDir) {
		t.Errorf("expected directory to not exist: %v", emptyDir)
	}
}

func TestRemoveEmptyDirNotEmpty(t *testing.T) {
	tempDir := t.TempDir()
	notEmptyDir := filepath.Join(tempDir, "not_empty_dir")

	// Create a directory with a file
	os.Mkdir(notEmptyDir, os.ModePerm)
	os.WriteFile(filepath.Join(notEmptyDir, "test.txt"), []byte("test"), 0644)

	// Test removing a non-empty directory
	err := RemoveEmptyDir(notEmptyDir)
	if err == nil {
		t.Errorf("expected error when removing non-empty directory: %v", notEmptyDir)
	}
}

func TestRemoveEmptyDirNonExistent(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent_dir")

	// Test removing a non-existent directory
	err := RemoveEmptyDir(nonExistentDir)
	if err == nil {
		t.Errorf("expected error when removing non-existent directory: %v", nonExistentDir)
	}
}

func TestEmptyDirW(t *testing.T) {
	tempDir := t.TempDir()
	dir := &Dir{Path: tempDir, Children: []string{"test"}}

	// Test emptying the directory
	emptyDirW(dir)

	// Verify directory is empty
	if dir.Path != "" || len(dir.Children) != 0 {
		t.Errorf("expected directory to be empty, got: %v", dir)
	}
}

func TestEmptyDirQ(t *testing.T) {
	tempDir := t.TempDir()
	dir := &Dir{Path: tempDir, Children: []string{"test"}}

	// Test emptying the directory
	children := emptyDirQ(dir)

	// Verify directory is empty
	if dir.Path != "" || len(dir.Children) != 0 {
		t.Errorf("expected directory to be empty, got: %v", dir)
	}

	// Verify children are returned
	if len(children) != 1 || children[0] != "test" {
		t.Errorf("expected children to be: %v, got: %v", []string{"test"}, children)
	}
}

func TestWalkFuncError(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test reading directory with error
	_, err := ReadDir(errDir)
	if err == nil {
		t.Errorf("expected error when reading directory with unreadable file: %v", errDir)
	}
}

func TestWalkFuncErrorReadDirQ(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test reading directory with error
	_, err := ReadDirQ(errDir)
	if err == nil {
		t.Errorf("expected error when reading directory with unreadable file: %v", errDir)
	}
}

func TestWalkFuncErrorReadDirW(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test reading directory with error
	err := ReadDirW(errDir)
	if err == nil {
		t.Errorf("expected error when reading directory with unreadable file: %v", errDir)
	}
}

func TestWalkFuncErrorReadDirA(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test reading directory with error
	d := &Dir{Path: errDir}
	err := ReadDirA(d)
	if err == nil {
		t.Errorf("expected error when reading directory with unreadable file: %v", errDir)
	}
}

func TestWalkFuncErrorReadDirD(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test reading directory with error
	id := ReadDirD(errDir)
	if id == "" {
		t.Errorf("expected to get a valid ID, got: %v", id)
	}
}

func TestRemoveDirQNonExistent(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent_dir")

	// Test removing a non-existent directory
	err := RemoveDirQ(nonExistentDir)
	if err == nil {
		t.Errorf("expected error when removing non-existent directory: %v", nonExistentDir)
	}
}

func TestRemoveDirWNonExistent(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent_dir")
	dir := &Dir{Path: nonExistentDir}

	// Test removing a non-existent directory
	err := RemoveDirW(dir)
	if err == nil {
		t.Errorf("expected error when removing non-existent directory: %v", nonExistentDir)
	}
}

func TestRemoveDirANonExistent(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent_dir")
	dir := &Dir{Path: nonExistentDir}

	// Test removing a non-existent directory
	_, err := RemoveDirA(dir)
	if err == nil {
		t.Errorf("expected error when removing non-existent directory: %v", nonExistentDir)
	}
}

func TestMoveDirDestinationExists(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "source")
	destinationDir := filepath.Join(tempDir, "destination")

	// Create both source and destination directories
	os.Mkdir(sourceDir, os.ModePerm)
	os.Mkdir(destinationDir, os.ModePerm)

	// Test moving directory with existing destination
	err := MoveDir(sourceDir, destinationDir)
	if err == nil {
		t.Errorf("expected error when moving directory to existing destination: %v", destinationDir)
	}
}

func TestMoveDirSourceNonExistent(t *testing.T) {
	tempDir := t.TempDir()
	sourceDir := filepath.Join(tempDir, "nonexistent_source")
	destinationDir := filepath.Join(tempDir, "destination")

	// Test moving a non-existent source directory
	err := MoveDir(sourceDir, destinationDir)
	if err == nil {
		t.Errorf("expected error when moving non-existent source directory: %v", sourceDir)
	}
}

func TestListFilesInDirNonExistent(t *testing.T) {
	tempDir := t.TempDir()
	nonExistentDir := filepath.Join(tempDir, "nonexistent_dir")

	// Test listing files in a non-existent directory
	_, err := ListFilesInDir(nonExistentDir)
	if err == nil {
		t.Errorf("expected error when listing files in non-existent directory: %v", nonExistentDir)
	}
}

func TestListFilesInDirError(t *testing.T) {
	tempDir := t.TempDir()
	errDir := filepath.Join(tempDir, "err_dir")

	// Create a directory with an error
	os.Mkdir(errDir, os.ModePerm)
	os.WriteFile(filepath.Join(errDir, "test.txt"), []byte("test"), 0644)
	os.Chmod(filepath.Join(errDir, "test.txt"), 0000) // Make file unreadable

	// Test listing files in directory with error
	_, err := ListFilesInDir(errDir)
	if err == nil {
		t.Errorf("expected error when listing files in directory with unreadable file: %v", errDir)
	}
}
