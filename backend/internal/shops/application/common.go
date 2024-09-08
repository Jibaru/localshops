package application

type ShopFieldsDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
}

type ShopDTO struct {
	ShopFieldsDTO
	ID string `json:"id"`
}
