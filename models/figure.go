package model

type Figure struct {
	Status       string `json:"status"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Series       string `json:"series"`
	Original     string `json:"original"`
	Scale        string `json:"scale"`
	Material     string `json:"material"`
	Price        string `json:"price(jpy)"`
	Code         string `json:"code"`
	JANCode      string `json:"jan_code"`
	ReleaseDate  string `json:"release_date"`
	Thumbnail    string `json:"thumbnail"`
	Source       string `json:"source"`
}

type FigureSearch struct {
	Keyword      string `query:"keyword"`
	ItemSeries   string `query:"item_series"`
	Manufacturer string `query:"manufacturer"`
	Status       string `query:"item_status"`
	Page         int    `query:"page"`
}
