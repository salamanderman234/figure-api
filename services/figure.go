package service

import (
	"salamanderman234/figure-price-api/data"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
)

type figureService struct {
	HobbySearchScrapper domain.FigureScrapper
	sourceMap           map[string]domain.FigureScrapper
}

func NewFigureService(h domain.FigureScrapper) domain.FigureService {
	sourceMap := map[string]domain.FigureScrapper{
		"hobby-search": h,
	}
	return &figureService{
		HobbySearchScrapper: h,
		sourceMap:           sourceMap,
	}
}

func (f *figureService) GetWithFilter(filter model.FigureSearch) ([]model.Figure, error) {

	// get source
	scrapper := f.sourceMap["hobby-search"]
	if filter.Source != "" {
		source, ok := f.sourceMap[filter.Source]
		if !ok {
			return nil, data.ErrNotFound
		}
		scrapper = source
	}

	if filter.ItemCode != "" {
		result, err := scrapper.DetailProduct(filter.ItemCode)
		if err != nil {
			return nil, err
		}
		return []model.Figure{result}, nil
	}
	result, err := scrapper.Search(filter)
	if err != nil {
		return nil, err
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
