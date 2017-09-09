package tempfiles

import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
)

func TestFileExistShouldFindExistingFile(T *testing.T) {
	folder := NewDirectory()
	file := folder.NewEmptyFile("Blob")
	defer folder.Remove()

	if !pathExists(string(file)) {
		T.Error("Folder should Exist")
	}

	file.ShouldExist(T, "folder wasn't deleted")
}

func TestRemovingFolderShouldAlsoRemoveTestfiles(T *testing.T) {
	folder := NewDirectory()
	file := folder.NewEmptyFile("Blob")
	folder.Remove()

	if pathExists(string(file)) {
		T.Error("File should not Exist")
	}

	file.ShouldNotExist(T, "folder was removed before")
}

func TestFilesShouldHaveSentContent(T *testing.T) {
	// arrange
	folder := NewDirectory()
	defer folder.Remove()
	path := filepath.Join(folder.Path(), "testfile")
	content := []byte{'a', '2', '3', '4'}

	// act
	file := folder.NewEmptyFile("Blob")
	ioutil.WriteFile(path, content, os.ModePerm)

	file.ShouldExist(T).WithContent([]byte{})
	folder.ShouldHaveFile(T, "testfile").WithContent(content, "this was the content we wrote to our new file")
}
