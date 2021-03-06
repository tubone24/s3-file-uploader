package router

import (
	"crypto/subtle"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/tubone24/s3-file-uploader/src/backend/config"
	"github.com/tubone24/s3-file-uploader/src/backend/logic"
	"github.com/tubone24/s3-file-uploader/src/backend/utils"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"strings"
)

var logicInstance = logic.New()

func New() *echo.Echo {
	appConfig := config.GetConfig()
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.POST},
	}))
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	e.Use(middleware.BasicAuthWithConfig(middleware.BasicAuthConfig{
		Skipper: func(c echo.Context) bool {
			if strings.HasSuffix(c.Request().RequestURI, "/status") {
				return true
			}
			return false
		},
		Validator: func(username, password string, c echo.Context) (bool, error) {
			if subtle.ConstantTimeCompare([]byte(username), []byte(appConfig.BasicAuth.UserName)) == 1 &&
				subtle.ConstantTimeCompare([]byte(password), []byte(appConfig.BasicAuth.Password)) == 1 {
				return true, nil
			}
			return false, nil
		},
		Realm: "Enter your credentials to login to log-uploader",
	}))
	initRouting(e)
	e.Validator = &CustomValidator{validator: validator.New()}
	return e
}

func initRouting(e *echo.Echo){
	e.Static("/", "assets")
	e.GET("/status", status)
	e.POST("/api/upload", upload)
	e.GET("/api/list", list)
	e.GET("/api/download", download)
	e.POST("/api/delete", deleteFile)
	e.GET("/api/checkdb", checkDB)
}

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}

func upload(c echo.Context) (err error) {
	data := new(UploadData)
	if err = c.Bind(data); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed bind data")
	}
	if err = c.Validate(data); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
	}
	result, err := logicInstance.UploadFileToS3(data.FileType, &data.Data, data.FileName)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, result)
	}
	return c.JSON(http.StatusOK, map[string]string{"fileType": data.FileType, "fileName": data.FileName})
}

func list(c echo.Context) (err error) {
	prefix := c.QueryParam("prefix")
	result, err := logicInstance.ListObjects(prefix)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	return c.JSON(http.StatusOK, &result)
}

func download(c echo.Context) (err error) {
	key := c.QueryParam("key")
	fileBytes, err := logicInstance.DownloadFileToS3(key)
	contentType := utils.GetContentType(key)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error")
	}
	return c.Blob(http.StatusOK, contentType, *fileBytes)
}

func deleteFile(c echo.Context) (err error) {
	data := new(DeletePayload)
	if err = c.Bind(data); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "failed bind data")
	}
	if err = c.Validate(data); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
	}

	err = logicInstance.DeleteFileToS3(data.Key)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error...")
	}
	return c.JSON(http.StatusOK, map[string]string{"fileName": data.Key})
}

func checkDB(c echo.Context) (err error) {
	ret, err := logicInstance.CheckDynamoLastKey()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "error...")
	}
	return c.JSON(http.StatusOK, ret)
}
