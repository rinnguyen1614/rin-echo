package upload

import "mime/multipart"

type Upload interface {
	Save(file *multipart.FileHeader) (*FileUploaded, error)
	Delete(key string) error
}

type FileUploaded struct {
	filename string
	path     string
	size     int64
	ext      string
}

func (f FileUploaded) Filename() string {
	return f.filename
}

func (f FileUploaded) Path() string {
	return f.path
}

func (f FileUploaded) Size() int64 {
	return f.size
}

func (f FileUploaded) Ext() string {
	return f.ext
}
