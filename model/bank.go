package model

// Province represents a province.
type Bank struct {
	ID        int    `json:"id" example:"1"`
	Bank      string `json:"bank" example:"BRI"`
	City      string `json:"city" example:"Surabaya"`
	Branch    string `json:"branch" example:"Surabaya"`
	SwiftCode string `json:"swift_code" example:"BBRIDJA"`
}

// ProvinceDataResponse represents the response format for the province data.
type BanksDataResponse struct {
	TotalCount int    `json:"total_count"` // Total number of provinces
	Banks      []Bank `json:"banks"`       // List of provinces
}

// Province represents a province.
type BankDataRequest struct {
	Bank      string `json:"bank" example:"BRI"`
	City      string `json:"city" example:"Surabaya"`
	Branch    string `json:"branch" example:"Surabaya"`
	SwiftCode string `json:"swift_code" example:"BBRIDJA"`
}

// Province represents a province.
type AddBanksDataResponse struct {
	Data    Bank   `json:"bank"`
	Message string `json:"message" example:"success create post"`
	Success bool   `json:"success" example:"true"`
}
