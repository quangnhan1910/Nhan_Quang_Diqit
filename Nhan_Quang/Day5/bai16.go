package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

var products []Product
var nextID = 1

func main() {
	r := gin.Default()

	r.POST("/products", createProduct)

	r.Run(":8080")
}

func createProduct(c *gin.Context) {
	var newProduct Product

	// 1. Bind JSON (Kiểm tra định dạng JSON)
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	// 2. Validate Tên (Không được rỗng)
	if newProduct.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tên sản phẩm không được để trống"})
		return
	}

	// 3. Validate Giá (Phải > 0)
	if newProduct.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Giá sản phẩm phải lớn hơn 0"})
		return
	}

	// 4. Validate Số lượng (Phải >= 0)
	if newProduct.Quantity < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Số lượng không được âm"})
		return
	}

	// 5. Nếu hợp lệ -> Lưu vào danh sách
	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	// Trả về 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Sản phẩm đã được tạo",
		"data":    newProduct,
	})
}
