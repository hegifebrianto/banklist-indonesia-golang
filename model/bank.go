package model

// Province represents a province.
type Bank struct {
	ID   int    `json:"id" example:"1"`
	Name string `json:"bankname" example:"BRI"`
}

// ProvinceDataResponse represents the response format for the province data.
type BanksDataResponse struct {
	TotalCount int    `json:"total_count"` // Total number of provinces
	Banks      []Bank `json:"banks"`       // List of provinces
}
