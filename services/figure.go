package service

import (
	"salamanderman234/figure-price-api/data"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
)

type figureService struct {
	HobbySearchScrapper domain.FigureScrapper
}

func NewFigureService(h domain.FigureScrapper) domain.FigureService {
	return &figureService{
		HobbySearchScrapper: h,
	}
}

func (f *figureService) GetWithFilter(filter model.FigureSearch) ([]model.Figure, error) {
	result, err := f.HobbySearchScrapper.Search(filter)

	if err != nil {
		return nil, err
	} else if len(result) == 0 {
		return nil, data.ErrNotFound
	}
	return result, nil
}
func (f *figureService) GetByShopCode(shopCode string) (model.Figure, error) {
	result, err := f.HobbySearchScrapper.DetailProduct(shopCode)
	if err != nil {
		return model.Figure{}, err
	}

	return result, nil
}
