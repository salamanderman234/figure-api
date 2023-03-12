package model

import "github.com/labstack/echo/v4"

type RouteList struct {
	Default *echo.Group
	Auth    *echo.Group
	Guest   *echo.Group
}
