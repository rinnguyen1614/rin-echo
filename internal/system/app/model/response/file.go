package response

import "github.com/rinnguyen1614/rin-echo/internal/core/utils/upload"

type FileResponse struct {
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Filename string `json:"file_name"`
	Ext      string `json:"ext"`
	Name     string `json:"name"`
}

func NewFile(file upload.FileUploaded) FileResponse {
	return FileResponse{
		Filename: file.Filename(),
		Path:     file.Path(),
		Size:     file.Size(),
		Ext:      file.Ext(),
		Name:     file.Name(),
	}
}
