package tests

import (
	"fmt"
	"github.com/aleksnew2/fs-utilsgo"
	"testing"
)

//func TestCreateFileQ(t *testing.T) {
//	file, ans := CreateFileQ("File1.txt")
//
//	fmt.Println(file)
//
//	if ans != nil {
//		t.Errorf("CreateFileQ fail: %v", ans)
//	}
//}

func TestCreateFileW(t *testing.T) {
	file, ans := fs_utils.CreateFileW("File2.txt", "BY")
	fmt.Println(file)
	if ans != nil {
		t.Errorf("CreateFileW fail: %v", ans)
	}
}

func TestGetFile(t *testing.T) {
	path, ans := fs_utils.GetFile("File2.txt")
	want := "File2.txt"
	fmt.Println(path)

	if path != want {
		t.Errorf("GetFile fail: %v", ans)
	}
}
