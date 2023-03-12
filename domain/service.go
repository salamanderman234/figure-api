package domain

import model "salamanderman234/figure-price-api/models"

type FigureService interface {
	GetWithFilter(filter model.FigureSearch) ([]model.Figure, error)
	GetByShopCode(shopCode string) (model.Figure, error)
}
