package tools

import (
	"encoding/json"
	"salamanderman234/figure-price-api/data"
	"salamanderman234/figure-price-api/domain"
	model "salamanderman234/figure-price-api/models"
	"strconv"
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
		searchUrl:        `/eng/search?`,
		detailProductUrl: `/eng/`,
	}
}

func (a *hobbySearch) DetailProduct(code string) (model.Figure, error) {
	// init
	figureMap := map[string]any{}
	figure := model.Figure{
		Series: "-",
	}
	// get
	resp, err := soup.Get(a.baseUrl + a.detailProductUrl + code)
	if err != nil {
		return model.Figure{}, err
	}
	doc := soup.HTMLParse(resp)
	// get information
	table := doc.Find("table", "id", "tblItemInfo")
	if table.Error == nil {
		datas := table.Find("tbody").FindAll("tr")
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
	} else {
		return model.Figure{}, data.ErrNotFound
	}
	// mapping to model format
	reserveInput := doc.Find("input", "class", "cart_button")
	if reserveInput.Error == nil {
		value := reserveInput.Attrs()["value"]
		mapValue := map[string]string{
			"Cart":    "Ready",
			"Reserve": "Pre-order",
		}
		figureMap["status"] = mapValue[value]
	} else {
		figureMap["status"] = "Sold Out"
	}
	price, ok := figureMap["sales_price"]
	if ok {
		figureMap["price(jpy)"] = price
		figureMap["sales_price"] = ""
	}
	name := doc.Find("h2", "class", "h2_itemDetail").Text()
	thumbnail := a.baseUrl
	thumbnailCandidate := doc.Find("table", "id", "masterBody_tblItemImg").Find("tbody").FindAll("img")
	if len(thumbnailCandidate) >= 3 {
		thumbnail += thumbnailCandidate[1].Attrs()["src"]
	}

	figureMap["name"] = name
	figureMap["thumbnail"] = thumbnail
	// mapping back to model
	jsonByte, _ := json.Marshal(figureMap)
	json.Unmarshal(jsonByte, &figure)

	return figure, nil
}
func (a *hobbySearch) Search(filter model.FigureSearch) ([]model.Figure, error) {
	// making filter
	page := "&spage=" + strconv.Itoa(filter.Page)
	keyword := "searchKey=" + strings.ReplaceAll(filter.Keyword, " ", "+")
	category := "&Typ2Cd=102"
	series := "&ItemSeries=" + strings.ReplaceAll(filter.ItemSeries, " ", "+")
	manufacturer := "&Make=" + strings.ReplaceAll(filter.Manufacturer, " ", "+")
	status := "&state="

	if filter.Status != "" {
		if filter.Status == "ready" {
			status += "4"
		} else if filter.Status == "preorder" {
			status += "1"
		}
	}
	// soup
	resp, err := soup.Get(a.baseUrl + a.searchUrl + keyword + category + series + manufacturer + status + page)
	if err != nil {
		return nil, err
	}
	doc := soup.HTMLParse(resp)
	// get relevant data
	figures := []model.Figure{}
	datas := doc.FindAll("div", "class", "ListItemName")
	if len(datas) > 0 {
		for _, data := range datas {
			href := data.Find("a").Attrs()["href"]
			kode := strings.ReplaceAll(href, "/eng/", "")
			figure, err := a.DetailProduct(kode)
			if err != nil {
				return nil, err
			}
			figures = append(figures, figure)
		}
	} else {
		return nil, data.ErrNotFound
	}
	return figures, nil
}
