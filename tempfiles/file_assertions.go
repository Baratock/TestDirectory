package tempfiles

import (
	"testing"
	"fmt"
)

func (file TestFile) ShouldExist(T *testing.T, args ...interface{}) {
	if !pathExists(file.Path()) {
		T.Error(fmt.Sprintf("File %s should exist", file), args)
	}
}

func (file TestFile) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(file.Path()) {
		T.Error(fmt.Sprintf("File %s should not exist", file), args)
	}
}
