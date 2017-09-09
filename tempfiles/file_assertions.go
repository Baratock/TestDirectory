package tempfiles

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"fmt"
)

type FileAssertion struct {
	t    *testing.T
	file TestFile
}

func (assertion *FileAssertion) WithContent(expected []byte, args ...interface{}) {
	if assertion == nil || assertion.file == "" {
		return
	}

	content, _ := ioutil.ReadFile(assertion.file.Path())
	assert.Equal(assertion.t, expected, content, args)
}

func (file TestFile) ShouldExist(T *testing.T, args ...interface{}) *FileAssertion {
	if !pathExists(file.Path()) {
		message := fmt.Sprintf("File %s should exist", file.Path())
		assert.Fail(T, message, args...)
		return nil
	}
	fa := &FileAssertion{
		T,
		file,
	}

	return fa
}

func (file TestFile) ShouldNotExist(T *testing.T, args ...interface{}) {
	if pathExists(file.Path()) {
		message := fmt.Sprintf("File %s should not exist", file.Path())
		assert.Fail(T, message, args...)
	}
}
