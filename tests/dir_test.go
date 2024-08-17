package tests

import (
	"fmt"
	"github.com/aleksnew2/fs-utilsgo"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := fs_utils.CreateDir("test1")
	if err != nil {
		t.Errorf("CreateDir fail: %v", err)
	}
}

func TestCreateDirQ(t *testing.T) {
	err := fs_utils.CreateDirQ("test1/test.txt")
	if err != nil {
		t.Errorf("CreateDirQ fail: %v", err)
	}
}

func TestCreateDirW(t *testing.T) {
	dir, err := fs_utils.CreateDirW("test1")
	fmt.Println(dir)
	if err != nil {
		t.Errorf("CreateDirW fail: %v", err)
	}
}

func TestReadDir(t *testing.T) {
	files, err := fs_utils.ReadDir("test1")
	fmt.Println(files)
	if err != nil {
		t.Errorf("ReadDir fail: %v", err)
	}
}

func TestReadDirW(t *testing.T) {
	err := fs_utils.ReadDirW("test1")
	if err != nil {
		t.Errorf("ReadDirW fail: %v", err)
	}
}

func TestReadDirQ(t *testing.T) {
	dir, err := fs_utils.ReadDirQ("test1")
	dir.Output()
	if err != nil {
		t.Errorf("ReadDirQ fail: %v", err)
	}
}
