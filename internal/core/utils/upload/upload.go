package upload

import (
	"mime/multipart"

	core "github.com/rinnguyen1614/rin-echo/internal/core"
)

var (
	ErrMaxUploadSize = core.NewRinError("max_upload_size", "uploaded file size exceeds the limit")
)

type Upload interface {
	Save(file *multipart.FileHeader, dst string) (*FileUploaded, error)
}

type FileUploaded struct {
	filename string
	path     string
	size     int64
	ext      string
	name     string
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

func (f FileUploaded) Name() string {
	return f.name
}
