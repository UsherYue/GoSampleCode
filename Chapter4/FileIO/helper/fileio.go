package helper

import (
	"io/ioutil"
	"os"
)

//创建FileIO对象
func NewFileIO(filename string) *FileIO {
	return &FileIO{filename: filename}
}

type FileIO struct {
	filename string
}

func (this *FileIO) CreateFile() error {
	file, err := os.Create(this.filename)
	defer file.Close()
	return err
}

func (this *FileIO) DeleteFile() error {
	return os.Remove(this.filename)
}

func (this *FileIO) WriteFile(data []byte) error {
	return ioutil.WriteFile(this.filename, data, 0777)
}

func (this *FileIO) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(this.filename)
}
