package tempfiles

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreatingFileShouldWork(T *testing.T) {
	// arrange
	dir := NewDirectory()
	defer dir.Remove()

	// act
	file1 := NewEmptyFile(dir, "NewFile.zip")
	file2 := NewEmptyFile(dir, "NewFile.xml")

	assert.Equal(T, "NewFile.zip", file1.Name())
	assert.Equal(T, "NewFile.xml", file2.Name())
}