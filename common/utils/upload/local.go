package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	filex "rin-echo/common/utils/file"
)

type Local struct {
}

func NewLocal() Upload {
	return &Local{}
}

func (l *Local) Save(file *multipart.FileHeader, dst string) (*FileUploaded, error) {
	var (
		dir, filename = path.Split(dst)
		ext           = path.Ext(dst)
	)

	// if dst parameter doesn't contain extention, set ext by file's filename then add extention for dst and filename
	if ext == "" {
		ext = path.Ext(file.Filename)
		dst += ext
		filename += ext
	}

	if err := filex.MkdirAll(dir); err != nil {
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return nil, err
	}

	return &FileUploaded{
		filename: filename,
		path:     dst,
		size:     file.Size,
		ext:      ext,
	}, nil
}
