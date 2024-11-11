package types

type Product struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Image       string   `json:"image"`
	Description string   `json:"description"`
	Category    Category `json:"category"`
}
