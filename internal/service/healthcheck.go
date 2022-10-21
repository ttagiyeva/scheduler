package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

//healthcheck default ping to server
func healthcheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}
