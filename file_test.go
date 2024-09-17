package fs_utils

import (
	"os"
	"testing"
)

func TestIsFileExists(t *testing.T) {
	// Test with a non-existing file
	if IsFileExists("non_existing_file.txt") {
		t.Error("Expected false for non-existing file")
	}

	// Create a temporary file for testing
	file, _ := os.Create("test_file.txt")
	file.Close()

	// Test with an existing file
	if !IsFileExists("test_file.txt") {
		t.Error("Expected true for existing file")
	}

	// Clean up
	os.Remove("test_file.txt")
}

func TestCreateFileQ(t *testing.T) {
	// Test creating a new file
	file, err := CreateFileQ("test_create_file.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if file.Path != "test_create_file.txt" {
		t.Error("File path does not match")
	}

	// Test creating a file that already exists
	_, err = CreateFileQ("test_create_file.txt")
	if err == nil {
		t.Error("Expected error for existing file")
	}

	// Clean up
	os.Remove("test_create_file.txt")
}

func TestCreateFileW(t *testing.T) {
	content := FileLines{"Line 1", "Line 2"}
	file, err := CreateFileW("test_create_file_w.txt", content)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if file.Path != "test_create_file_w.txt" {
		t.Error("File path does not match")
	}

	// Check file content
	lines, _ := GetFileContent("test_create_file_w.txt")
	if len(lines) != 2 {
		t.Error("Expected 2 lines in the file")
	}

	// Clean up
	os.Remove("test_create_file_w.txt")
}

func TestRemoveFileQ(t *testing.T) {
	// Create a file to remove
	os.Create("test_remove_file_q.txt")

	// Test removing the file
	err := RemoveFileQ("test_remove_file_q.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test removing a non-existing file
	err = RemoveFileQ("test_remove_file_q.txt")
	if err == nil {
		t.Error("Expected error for non-existing file")
	}
}

func TestRenameFile(t *testing.T) {
	// Create a file to rename
	os.Create("test_rename_file.txt")

	// Test renaming the file
	err := RenameFile("test_rename_file.txt", "test_renamed_file.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the new file exists
	if !IsFileExists("test_renamed_file.txt") {
		t.Error("Expected renamed file to exist")
	}

	// Clean up
	os.Remove("test_renamed_file.txt")
}

func TestCopyFile(t *testing.T) {
	// Create a source file
	sourceFile, _ := os.Create("test_copy_source.txt")
	sourceFile.WriteString("This is a test.")
	sourceFile.Close()

	// Test copying the file
	err := CopyFile("test_copy_source.txt", "test_copy_destination.txt")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the destination file exists
	if !IsFileExists("test_copy_destination.txt") {
		t.Error("Expected copied file to exist")
	}

	// Clean up
	os.Remove("test_copy_source.txt")
	os.Remove("test_copy_destination.txt")
}

func TestAppendToFile(t *testing.T) {
	// Create a file to append to
	os.Create("test_append_file.txt")

	content := FileLines{"Line 1", "Line 2"}
	err := AppendToFile("test_append_file.txt", content)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check file content
	lines, _ := GetFileContent("test_append_file.txt")
	if len(lines) != 2 {
		t.Error("Expected 2 lines in the file after append")
	}

	// Clean up
	os.Remove("test_append_file.txt")
}