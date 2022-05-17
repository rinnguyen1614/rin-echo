package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"
	"rin-echo/common/utils"
	"strings"
)

type Local struct {
	basePath     string
	hashFilename bool
	format       string
}

func NewLocal(basePath string) *Local {
	return &Local{
		basePath: path.Clean(basePath),
	}
}

func (l *Local) Save(file *multipart.FileHeader) (*FileUploaded, error) {
	var (
		filename           = file.Filename
		ext                = path.Ext(filename)
		filenameWithoutExt = strings.TrimSuffix(filename, ext)
		dst                string
	)

	if l.hashFilename {
		filename = utils.MD5([]byte(filenameWithoutExt)) + ext
	}
	dst = l.basePath + "/" + filename

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
