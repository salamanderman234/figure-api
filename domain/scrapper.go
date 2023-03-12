package domain

import model "salamanderman234/figure-price-api/models"

type FigureScrapper interface {
	Search(filter model.FigureSearch) ([]model.Figure, error)
	DetailProduct(code string) (model.Figure, error)
}
