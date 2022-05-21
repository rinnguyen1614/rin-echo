package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	filex "rin-echo/common/utils/file"
	"strings"
)

const MAX_UPLOAD_SIZE = filex.MB * 5

type Local struct {
	// MaxUploadSize limits the size of the uploaded content, specified with "byte".
	MaxUploadSize int64
}

func NewLocal(maxUploadSize int64) Upload {
	if maxUploadSize <= 0 {
		maxUploadSize = MAX_UPLOAD_SIZE
	}
	return &Local{MaxUploadSize: maxUploadSize}
}

func (l *Local) Save(file *multipart.FileHeader, dst string) (*FileUploaded, error) {
	var (
		dir, filename = path.Split(dst)
		ext           = path.Ext(dst)
		size          = file.Size
	)

	if size > l.MaxUploadSize {
		return nil, ErrMaxUploadSize
	}

	// if dst parameter doesn't contain extention, the ext is assigned by ext of file.filename then add extention for the dst and the filename
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

	if !strings.HasPrefix(dst, "/") {
		dst = "/" + dst
	}

	return &FileUploaded{
		filename: filename,
		path:     dst,
		size:     size,
		ext:      ext,
		name:     file.Filename,
	}, nil
}
