package domain

import "github.com/labstack/echo/v4"

type FigureController interface {
	SearchFiguresWithFilter(c echo.Context) error
	SearchFigureWithShopCode(c echo.Context) error
}

type ViewController interface {
	Index(c echo.Context) error
}
