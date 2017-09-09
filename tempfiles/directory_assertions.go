package tempfiles

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"fmt"
)

func (directory TestDirectory) ShouldExist(T *testing.T, args ...interface{}) {
	if !pathExists(directory.Path()) {
		message := fmt.Sprintf("Directory %s should exist", directory.Path())
		assert.Fail(T, message, args...)
	}
}

func (directory TestDirectory) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(directory.Path()) {
		message := fmt.Sprintf("Directory %s should not exist", directory.Path())
		assert.Fail(T, message, args)
	}
}

func pathExists(directory string) bool {
	_, err := os.Stat(directory)

	if os.IsNotExist(err) {
		return false;
	}
	return true
}

func (directory TestDirectory) ShouldHaveFile(T *testing.T, filename string, args ...interface{}) *FileAssertion {
	path := filepath.Join(string(directory), filename)
	file := TestFile(path)

	return file.ShouldExist(T, args...)
}
