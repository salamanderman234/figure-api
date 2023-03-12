package controller

import (
	"errors"
	"net/http"
	"salamanderman234/figure-price-api/data"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"

	"github.com/labstack/echo/v4"
)

type figureController struct {
	figureService domain.FigureService
}

func NewFigureController(f domain.FigureService) domain.FigureController {
	return &figureController{
		figureService: f,
	}
}

func (f *figureController) SearchFiguresWithFilter(c echo.Context) error {
	// get relevant data from request
	filter := model.FigureSearch{}
	err := (&echo.DefaultBinder{}).BindQueryParams(c, filter)
	if err != nil {
		response := model.FailResponse{
			Error:   "Bad Request",
			Message: "Missing required field",
		}
		return c.JSON(http.StatusBadRequest, response)
	}
	// calling service
	result, err := f.figureService.GetWithFilter(filter)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorType := "Internal server error"
		errorMessage := "something went wrong"
		if errors.Is(err, data.ErrNotFound) {
			statusCode = http.StatusNotFound
			errorType = "Not found"
			errorMessage = "Nothing found with that filter"
		}

		return c.JSON(statusCode, model.FailResponse{Error: errorType, Message: errorMessage})
	}
	// creating response
	response := model.SuccessResponse{
		Found: len(result),
		Data:  result,
	}
	return c.JSON(http.StatusOK, response)
}
func (f *figureController) SearchFigureWithShopCode(c echo.Context) error {
	// get relevant data from request
	shopCode := c.Param("shop_code")
	// calling service
	result, err := f.figureService.GetByShopCode(shopCode)
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorType := "Internal server error"
		errorMessage := "something went wrong"
		if errors.Is(err, data.ErrNotFound) {
			statusCode = http.StatusNotFound
			errorType = "Not found"
			errorMessage = "Nothing found with that filter"
		}

		return c.JSON(statusCode, model.FailResponse{Error: errorType, Message: errorMessage})
	}
	// creating response
	response := model.SuccessResponse{
		Found: 1,
		Data:  []model.Figure{result},
	}
	return c.JSON(http.StatusOK, response)
}
