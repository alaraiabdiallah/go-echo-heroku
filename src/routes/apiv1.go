package routes

import (
	"net/http"

	"github.com/labstack/echo"
)

func APIV1(api *echo.Group) {
	v1 := api.Group("v1")
	v1.GET("/", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Hello World")
	})
}
