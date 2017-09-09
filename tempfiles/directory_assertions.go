package tempfiles

import (
	"testing"
	"fmt"
	"os"
)

func (directory TestDirectory) ShouldExist(T *testing.T, args ...interface{}) {
	if !pathExists(directory.Path()) {
		T.Error(fmt.Sprintf("File %s should exist", directory), args)
	}
}

func (directory TestDirectory) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(directory.Path()) {
		T.Error(fmt.Sprintf("File %s should not exist", directory), args)
	}
}

func pathExists(directory string) bool {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		return false;
	}
	return true
}