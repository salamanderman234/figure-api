package model

type Figure struct {
	Code         string   `json:"item_code"`
	JANCode      string   `json:"jan_code"`
	Name         string   `json:"name"`
	Manufacturer string   `json:"manufacturer"`
	Series       string   `json:"series"`
	Original     string   `json:"original"`
	Scale        string   `json:"scale"`
	Material     string   `json:"material"`
	Price        Currency `json:"price"`
	ReleaseDate  string   `json:"release_date"`
	Thumbnail    string   `json:"thumbnail"`
}

type FigureSearch struct {
	Keyword      string `query:"keyword"`
	ItemSeries   string `query:"item_series"`
	Manufacturer string `query:"manufacturer"`
	Status       string `query:"item_status"`
}
