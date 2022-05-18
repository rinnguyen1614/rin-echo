package response

import "rin-echo/common/utils/upload"

type FileResponse struct {
	Path     string `json:"path"`
	Size     int64  `json:"size"`
	Filename string `json:"file_name"`
	Ext      string `json:"ext"`
}

func NewFile(file upload.FileUploaded) FileResponse {
	return FileResponse{
		Filename: file.Filename(),
		Path:     file.Path(),
		Size:     file.Size(),
		Ext:      file.Ext(),
	}
}
