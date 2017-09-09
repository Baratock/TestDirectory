package tempfiles

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"fmt"
	"testing"
)

type TestDirectory string

func (directory TestDirectory) ShouldHaveFile(T *testing.T, filename string, args ...interface{}) {
	path := filepath.Join(string(directory), filename)
	file := TestFile(path)

	file.ShouldExist(T, args)
}

func (dir TestDirectory) Remove() {
	err := os.RemoveAll(string(dir))
	if err != nil {
		fmt.Println(err)
	}
}

func NewDirectory() TestDirectory {
	dir, err := ioutil.TempDir("", "golangTest_")
	if err != nil {
		panic(err)
	}
	return TestDirectory(dir);
}

func (directory TestDirectory) Name() string {
	name := filepath.Base(directory.Path())
	return name
}

func (directory TestDirectory) Path() string {
	return string(directory)
}

func (directory TestDirectory) AddDirectory(name string) TestDirectory {
	path := filepath.Join(directory.Path(), name)
	err := os.Mkdir(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
	return TestDirectory(path);
}
