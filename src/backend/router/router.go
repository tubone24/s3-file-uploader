package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/logic"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
	}))
	initRouting(e)
	e.Validator = &CustomValidator{validator: validator.New()}
	return e
}

func initRouting(e *echo.Echo){
	e.Static("/", "assets")
	e.GET("/status", status)
	e.POST("/upload", upload)
}

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func upload(c echo.Context) (err error) {
	data := new(Data)
	if err = c.Bind(data); err != nil {
		return err
	}
	if err = c.Validate(data); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "BadRequest", "message": "invalid payload"})
	}
	result, err := logic.UploadFileToS3(data.FileType, data.Data, data.FileName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "InternalServerError", "message": result})
	}
	return c.JSON(http.StatusOK, map[string]string{"fileType": data.FileType, "fileName": data.FileName})
}
