package tests

import (
	"github.com/aleksnew2/fs-utilsgo"
	"testing"
)

func TestCreateDir(t *testing.T) {
	err := fs_utils.CreateDir("testdir")
	if err != nil {
		t.Errorf("CreateDir: %v", err)
	}
}
