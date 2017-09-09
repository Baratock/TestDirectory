package tempfiles

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
)

func (directory TestDirectory) ShouldExist(T *testing.T, args ...interface{}) {
	if !pathExists(directory.Path()) {
		assert.Fail(T, "Directory %s should exist", directory.Path(), args)
	}
}

func (directory TestDirectory) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(directory.Path()) {
		assert.Fail(T, "Directory %s should not exist", directory.Path(), args)
	}
}

func pathExists(directory string) bool {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		return false;
	}
	return true
}
