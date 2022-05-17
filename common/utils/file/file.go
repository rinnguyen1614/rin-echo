package file

import (
	"bytes"
	"os"
)

func Create(content bytes.Buffer, name string) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}

	_, err = file.WriteString(content.String())
	if err != nil {
		return err
	}

	file.Close()
	return nil
}

func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Size(path string) (int64, error) {
	file, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
}

func MkdirAll(path string) error {
	if !IsExist(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}
