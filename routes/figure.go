package route

import (
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
)

type figureRoute struct {
	figureController domain.FigureController
	groups           model.RouteList
	basePath         string
}

func NewFigureRoute(f domain.FigureController, groupList model.RouteList) domain.Route {
	return &figureRoute{
		figureController: f,
		groups:           groupList,
		basePath:         "figure",
	}
}

func (f *figureRoute) RegisterRoutes() {
	mainRoute := f.groups.Default
	mainRoute.GET(f.basePath, f.figureController.SearchFiguresWithFilter)
	mainRoute.GET(f.basePath+"/:shop_code", f.figureController.SearchFigureWithShopCode)
}

// route list
// figure -> all list with filter
// figure/:shop_code -> searching spesific figure base on shop_code
