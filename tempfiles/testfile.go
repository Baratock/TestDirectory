package tempfiles

import (
	"os"
	"path/filepath"
)

type TestFile string

func (file TestFile) Name() string {
	name := filepath.Base(file.Path())
	return name
}

func (file TestFile) Path() string {
	return string(file)
}

func NewEmptyFile(path TestDirectory, name string) TestFile {
	file, err :=os.Create(filepath.Join(path.Path(), name))

	if err != nil {
		panic(err)
	}

	return TestFile(file.Name())
}