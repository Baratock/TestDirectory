package tempfiles

import "testing"

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