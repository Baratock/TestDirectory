package tempfiles

import (
	"testing"
	"os"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/assert"
)

func TestThatTestDirectoryWasCreated(T *testing.T) {
	// creation
	// arrange & act
	dir := NewDirectory()
	var name = dir.Path()
	_, err := os.Stat(name)

	// assert
	require.Equal(T, nil, err, "Tempfolder was createtd")

	// deletion
	// act
	dir.Remove()
	_, err = os.Stat(name)
	DoesNotExist := os.IsNotExist(err)
	assert.Equal(T, true, DoesNotExist)
}

func TestAddedSubfolderShouldExist(T *testing.T) {
	// arrange
	dir := NewDirectory()
	defer dir.Remove()

	// act
	subdir1 := dir.AddDirectory("Dir1")
	subdir2 := dir.AddDirectory("Dir2")

	// assert
	assert.NotEqual(T, subdir1, subdir2)
	assert.Equal(T, "Dir1", subdir1.Name())
	assert.Equal(T, "Dir2", subdir2.Name())
}

func TestCreationOfFileFromDirShouldWork(T *testing.T) {
	dir := NewDirectory()
	defer dir.Remove()

	// act
	file1 := dir.NewEmptyFile("NewFile.zip")
	file2 := dir.NewEmptyFile("NewFile.xml")

	assert.Equal(T, "NewFile.zip", file1.Name())
	assert.Equal(T, "NewFile.xml", file2.Name())
}

func (directory TestDirectory) NewEmptyFile(name string) TestFile {
	file := NewEmptyFile(directory, name)
	return file
}
