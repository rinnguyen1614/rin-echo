package handler

import (
	"rin-echo/system/app/model/response"
	"rin-echo/system/domain"
	"rin-echo/system/inject"
	"strings"
	"time"

	core "github.com/rinnguyen1614/rin-echo-core"
	echox "github.com/rinnguyen1614/rin-echo-core/echo"
	rquery "github.com/rinnguyen1614/rin-echo-core/echo/models/query/rest_query"
	"github.com/rinnguyen1614/rin-echo-core/setting"
	iuow "github.com/rinnguyen1614/rin-echo-core/uow/interfaces"
	"github.com/rinnguyen1614/rin-echo-core/utils"
	"github.com/rinnguyen1614/rin-echo-core/utils/upload"
	"github.com/rinnguyen1614/rin-echo-core/validation"

	"go.uber.org/zap"
)

var (
	ErrPathParamRequired = core.NewRinError("path_param_required", "path_param_required")
)

type FileHandler struct {
	echox.Handler
	upload upload.Upload
}

func NewFileHandler(uow iuow.UnitOfWork,
	permissionManager domain.PermissionManager,
	logger *zap.Logger,
	restQuery rquery.RestQuery,
	settingProvider setting.Provider,
	validator *validation.Validator) FileHandler {

	uploadMaxSize := setting.MustGet[int64](settingProvider, "files.upload.max_size")
	return FileHandler{
		Handler: echox.NewHandler(logger, restQuery, settingProvider, validator),
		upload:  upload.NewLocal(uploadMaxSize),
	}
}

// UploadFile godoc
// @Summary      Upload file
// @Description  Upload file with the input payload
// @Tags         files
// @Accept       multipart/form-data
// @Produce      application/json
// @Param 		 files formData file true "Muilti files"
// @Success      200  {object}  models.Response{data=[]response.FileResponse} "{"data": {}}"
// @Router       /files/upload [post]
// @Security Bearer
func (h FileHandler) Upload(c echox.Context) error {
	session := c.MustSession().(*inject.Claims)
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]
	fileUploadeds := make([]response.FileResponse, 0)
	basePath, err := h.SettingProvider.Get("files.upload.path")

	if err != nil {
		return err
	}

	for _, file := range files {
		path := basePath + "/" + utils.Encrypt(session.Username, time.Now().Format("2006010250101")+utils.RandomLetter(2))
		fileUploaded, err := h.upload.Save(file, path)
		if err != nil {
			return err
		}
		fileUploadeds = append(fileUploadeds, response.NewFile(*fileUploaded))
	}

	echox.OKWithData(c, fileUploadeds)
	return nil
}

// DownloadFile godoc
// @Summary      Download file
// @Description  Download file with query's path
// @Tags         files
// @Param 		 path query string true "path of file"
// @Success      200
// @Router       /files/download [post]
// @Security Bearer
func (h FileHandler) Download(c echox.Context) error {
	path := c.QueryParam("path")
	if path == "" {
		return ErrPathParamRequired
	}
	// validator, ok := c.Echo().Validator.(*validation.Validator)
	// if ok {
	// 	if err := validator.Instance().Var(path, "required,url"); err != nil {
	// 		return err
	// 	}
	// }
	return c.File(strings.TrimPrefix(path, "/"))
}
