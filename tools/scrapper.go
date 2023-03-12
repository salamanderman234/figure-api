package tools

import (
	"encoding/json"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
	"strings"

	"github.com/anaskhan96/soup"
)

type hobbySearch struct {
	baseUrl          string
	searchUrl        string
	detailProductUrl string
}

func NewHobbySearchScrapper() domain.FigureScrapper {
	base := `https://www.1999.co.jp`
	return &hobbySearch{
		baseUrl:          base,
		searchUrl:        `/eng/search?Typ2Cd=102&searchkey=kato+megumi`,
		detailProductUrl: `/eng/`,
	}
}

func (a *hobbySearch) DetailProduct(code string) (model.Figure, error) {
	// init
	figureMap := map[string]any{}
	figure := model.Figure{}
	// get
	resp, _ := soup.Get(a.baseUrl + a.detailProductUrl + code)
	doc := soup.HTMLParse(resp)
	// get information
	datas := doc.Find("table", "id", "tblItemInfo").Find("tbody").FindAll("tr")
	for _, data := range datas {
		labelAndValues := data.FindAll("td")
		if len(labelAndValues) == 3 {
			key := strings.ReplaceAll(strings.ToLower(labelAndValues[0].Text()), " ", "_")
			valueContainer := labelAndValues[2]
			values := valueContainer.Find("a")
			if values.Error == nil {
				figureMap[key] = values.Text()
			} else {
				span := valueContainer.Find("span")
				if span.Error == nil {
					secondSpan := span.Find("span")
					if secondSpan.Error == nil {
						figureMap[key] = secondSpan.Text()
					} else {
						figureMap[key] = span.Text()
					}

				} else {
					figureMap[key] = valueContainer.Text()
				}

			}
		}

	}

	// mapping to model format
	price, ok := figureMap["sales_price"]
	if ok {
		figureMap["price"] = model.Currency{
			JPY: price.(string),
		}
		figureMap["sales_price"] = ""
	}
	name := doc.Find("h2", "class", "h2_itemDetail").Text()
	thumbnail := a.baseUrl
	thumbnailCandidate := doc.Find("table", "id", "masterBody_tblItemImg").Find("tbody").FindAll("img")
	if len(thumbnailCandidate) >= 3 {
		thumbnail += thumbnailCandidate[2].Attrs()["src"]
	}

	figureMap["name"] = name
	figureMap["thumbnail"] = thumbnail
	// mapping back to model
	jsonByte, _ := json.Marshal(figureMap)
	json.Unmarshal(jsonByte, &figure)

	return figure, nil
}
func (a *hobbySearch) Search(filter model.FigureSearch) ([]model.Figure, error) {
	return nil, nil
}
