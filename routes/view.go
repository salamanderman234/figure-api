package route

import (
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
)

type viewRoute struct {
	viewController domain.ViewController
	groups         model.RouteList
	basePath       string
}

func NewViewRoute(v domain.ViewController, groupList model.RouteList) domain.Route {
	return &viewRoute{
		viewController: v,
		groups:         groupList,
		basePath:       "",
	}
}

func (v *viewRoute) RegisterRoutes() {
	mainRoute := v.groups.Default
	mainRoute.GET("", v.viewController.Index)
}
