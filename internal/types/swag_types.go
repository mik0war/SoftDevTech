package types

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type OrderData struct {
	Products []struct {
		ProductId int `json:"product_id"`
		Count     int `json:"count"`
	} `json:"products"`
}

type ProductData struct {
	Name           string `json:"name"`
	Status         string `json:"status"`
	Description    string `json:"description"`
	ImageLink      string `json:"imageLink"`
	ManufacturerId int    `json:"manufacturer_id"`
}
