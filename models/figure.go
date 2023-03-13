package model

type Figure struct {
	Source       string `json:"source"`
	Code         string `json:"code"`
	Status       string `json:"status"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Series       string `json:"series"`
	Original     string `json:"original"`
	Scale        string `json:"scale"`
	Height       string `json:"height"`
	Material     string `json:"material"`
	Price        int    `json:"price(jpy)"`
	Description  string `json:"description"`
	// JANCode      string `json:"jan_code"`
	ReleaseDate string `json:"release_date"`
	Thumbnail   string `json:"thumbnail"`
}

type FigureSearch struct {
	Source       string `query:"source"`
	Keyword      string `query:"keyword"`
	ItemSeries   string `query:"item_series"`
	Manufacturer string `query:"manufacturer"`
	Status       string `query:"item_status"`
	Page         int    `query:"page"`
	ItemCode     string `query:"item_code"`
}
