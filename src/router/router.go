package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	//"github.com/tubone24/s3-file-uploader/src/logic"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	initRouting(e)
	e.Validator = NewValidator()
	return e
}

func initRouting(e *echo.Echo){
	e.Static("/", "assets")
	e.GET("/status", status)
}

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
