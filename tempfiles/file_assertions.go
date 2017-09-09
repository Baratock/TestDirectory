package tempfiles

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func (file TestFile) ShouldExist(T *testing.T, args ...interface{}) {
	if !pathExists(file.Path()) {
		assert.Fail(T, "File %s should exist", file.Path(), args)
	}
}

func (file TestFile) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(file.Path()) {
		assert.Fail(T, "File %s should not exist", file.Path(), args)
	}
}
