package valobj

import (
	"errors"
	"os"
	"path"
)

type File struct {
	path  string
	bytes []byte
}

func FileFromPath(path string) *File {
	return &File{path: path}
}

func SaveFileFromBytes(bytes []byte, filename string) (*File, error) {
	f := &File{bytes: bytes}

	f.setPath(filename)

	err := f.Save()

	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *File) setPath(filename string) {
	basePath := os.Getenv("FILE_SOTRAGE_PATH")
	f.path = path.Join(basePath, filename)
}

func (f *File) Load() error {
	bytes, err := os.ReadFile(f.path)

	if err != nil {
		return err
	}

	f.bytes = bytes

	return nil
}

func (f *File) Clean() {
	f.bytes = nil
}

func (f *File) Save() error {
	dir, _ := path.Split(f.path)

	_, err := os.Stat(dir)

	if errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dir, os.ModePerm)

		if err != nil {
			return err
		}
	}

	err = os.WriteFile(f.path, f.bytes, os.ModePerm)

	return err
}

func (f *File) String() string {
	return f.path
}
