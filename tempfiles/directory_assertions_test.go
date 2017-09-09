package tempfiles

import (
	"testing"
)

func TestFolderExistShouldFindExistingFolder(T *testing.T) {
	folder := NewDirectory()
	defer folder.Remove()

	if !pathExists(string(folder)) {
		T.Error("Folder should Exist")
	}

	folder.ShouldExist(T, "folder wasn't deleted")
}

func TestFolderExistShouldNotFindUnexistingFolder(T *testing.T) {
	// arrange
	folder := NewDirectory()
	folder.Remove()

	// assert
	if pathExists(string(folder)) {
		T.Error("Folder should not Exist")
	}

	folder.ShouldNotExist(T, "folder was removed before")
}
