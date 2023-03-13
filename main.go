package main

import (
	controller "salamanderman234/figure-price-api/controllers"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
	route "salamanderman234/figure-price-api/routes"
	service "salamanderman234/figure-price-api/services"
	"salamanderman234/figure-price-api/tools"

	"github.com/labstack/echo/v4"
)

func main() {
	// init
	router := echo.New()
	groupList := model.RouteList{}
	apiGroupList := model.RouteList{}
	routeList := []domain.Route{}
	// creating route group
	defaultGroup := router.Group("/")
	defaultApiGroup := router.Group("/api/")
	groupList.Default = defaultGroup
	apiGroupList.Default = defaultApiGroup
	// scrapper
	amiAmiScrapper := tools.NewHobbySearchScrapper()
	// service
	figureService := service.NewFigureService(amiAmiScrapper)

	// controller
	figureController := controller.NewFigureController(figureService)

	// route
	routeList = append(routeList, route.NewFigureRoute(figureController, apiGroupList))

	// register route
	for _, route := range routeList {
		route.RegisterRoutes()
	}
	// start
	router.Start(":1323")
}
