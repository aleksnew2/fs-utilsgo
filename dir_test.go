package fs_utils

import (
	"os"
	"testing"
)

func TestIsDirExists(t *testing.T) {
	// Test with a non-existing directory
	if IsDirExists("non_existing_dir") {
		t.Error("Expected false for non-existing directory")
	}

	// Create a temporary directory for testing
	os.Mkdir("test_dir", os.ModePerm)
	defer os.RemoveAll("test_dir") // Clean up

	// Test with an existing directory
	if !IsDirExists("test_dir") {
		t.Error("Expected true for existing directory")
	}
}

func TestGetDir(t *testing.T) {
	// Create a temporary directory
	os.Mkdir("test_get_dir", os.ModePerm)
	defer os.RemoveAll("test_get_dir") // Clean up

	// Test getting an existing directory
	dir, err := GetDir("test_get_dir")
	if err != nil || dir != "test_get_dir" {
		t.Errorf("Expected 'test_get_dir', got %v, error: %v", dir, err)
	}

	// Test getting a non-existing directory
	_, err = GetDir("non_existing_dir")
	if err == nil {
		t.Error("Expected error for non-existing directory")
	}
}

func TestGetDirQ(t *testing.T) {
	// Create a temporary directory with a file
	os.Mkdir("test_get_dir_q", os.ModePerm)
	os.Create("test_get_dir_q/test_file.txt")
	defer os.RemoveAll("test_get_dir_q") // Clean up

	// Test getting directory and its children
	dir, err := GetDirQ(&Dir{Path: "test_get_dir_q"})
	if err != nil || dir.Path != "test_get_dir_q" || len(dir.Children) != 1 {
		t.Errorf("Expected directory with 1 child, got %v, error: %v", dir, err)
	}
}

func TestReadDir(t *testing.T) {
	// Create a temporary directory with files
	os.Mkdir("test_read_dir", os.ModePerm)
	os.Create("test_read_dir/file1.txt")
	os.Create("test_read_dir/file2.txt")
	defer os.RemoveAll("test_read_dir") // Clean up

	// Test reading the directory
	files, err := ReadDir("test_read_dir")
	if err != nil || len(files) != 2 {
		t.Errorf("Expected 2 files, got %v, error: %v", files, err)
	}
}

func TestReadDirQ(t *testing.T) {
	// Create a temporary directory with files
	os.Mkdir("test_read_dir_q", os.ModePerm)
	os.Create("test_read_dir_q/file1.txt")
	defer os.RemoveAll("test_read_dir_q") // Clean up

	// Test reading the directory into a Dir object
	dir, err := ReadDirQ("test_read_dir_q")
	if err != nil || len(dir.Children) != 1 {
		t.Errorf("Expected 1 child, got %v, error: %v", dir, err)
	}
}

func TestReadDirW(t *testing.T) {
	// Create a temporary directory with files
	os.Mkdir("test_read_dir_w", os.ModePerm)
	os.Create("test_read_dir_w/file1.txt")
	defer os.RemoveAll("test_read_dir_w") // Clean up

	// Test reading the directory and printing output
	err := ReadDirW("test_read_dir_w")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestCreateDir(t *testing.T) {
	// Test creating a new directory
	err := CreateDir("test_create_dir")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.RemoveAll("test_create_dir") // Clean up

	// Test creating a directory that already exists
	err = CreateDir("test_create_dir")
	if err == nil {
		t.Error("Expected error for existing directory")
	}
}

func TestRemoveDirQ(t *testing.T) {
	// Create a temporary directory
	os.Mkdir("test_remove_dir_q", os.ModePerm)
	defer os.RemoveAll("test_remove_dir_q") // Clean up

	// Test removing the directory
	err := RemoveDirQ("test_remove_dir_q")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test removing a non-existing directory
	err = RemoveDirQ("test_remove_dir_q")
	if err == nil {
		t.Error("Expected error for non-existing directory")
	}
}

func TestMoveDir(t *testing.T) {
	// Create a temporary directory
	os.Mkdir("test_move_dir", os.ModePerm)
	defer os.RemoveAll("test_move_dir") // Clean up

	// Test moving the directory
	err := MoveDir("test_move_dir", "test_move_dir_new")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	defer os.RemoveAll("test_move_dir_new") // Clean up

	// Test moving to an existing directory
	os.Mkdir("test_move_dir_existing", os.ModePerm)
	defer os.RemoveAll("test_move_dir_existing") // Clean up
	err = MoveDir("test_move_dir_new", "test_move_dir_existing")
	if err == nil {
		t.Error("Expected error for moving to existing directory")
	}
}

func TestListFilesInDir(t *testing.T) {
	// Create a temporary directory with files
	os.Mkdir("test_list_files", os.ModePerm)
	os.Create("test_list_files/file1.txt")
	defer os.RemoveAll("test_list_files") // Clean up

	// Test listing files in the directory
	files, err := ListFilesInDir("test_list_files")
	if err != nil || len(files) != 1 {
		t.Errorf("Expected 1 file, got %v, error: %v", files, err)
	}
}

func TestRemoveEmptyDir(t *testing.T) {
	// Create a temporary directory
	os.Mkdir("test_remove_empty_dir", os.ModePerm)
	defer os.RemoveAll("test_remove_empty_dir") // Clean up

	// Test removing the empty directory
	err := RemoveEmptyDir("test_remove_empty_dir")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test removing a non-existing directory
	err = RemoveEmptyDir("test_remove_empty_dir")
	if err == nil {
		t.Error("Expected error for non-existing directory")
	}
}