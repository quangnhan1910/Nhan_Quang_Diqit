package models

// Chỉ chứa định nghĩa cấu trúc dữ liệu
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}