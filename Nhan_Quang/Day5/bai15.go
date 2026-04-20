// bai15.go - CRUD Sản phẩm với Gin Framework
package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description,omitempty"`
}

// Dữ liệu lưu tạm
var products []Product
var nextID = 4

func main() {
	// Dữ liệu mẫu ban đầu
	products = []Product{
		{ID: 1, Name: "iPhone 15", Price: 24990000, Description: "Điện thoại Apple"},
		{ID: 2, Name: "Samsung Galaxy S24", Price: 21990000, Description: "Điện thoại Samsung"},
		{ID: 3, Name: "Laptop Dell XPS 13", Price: 32990000, Description: "Laptop cao cấp"},
	}

	r := gin.Default()

	// CRUD Routes
	r.GET("/products", getAllProducts)
	r.GET("/products/:id", getProductByID)
	r.POST("/products", createProduct)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	r.Run(":8080")
}

// ==================== HANDLERS ====================

func getAllProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"count":  len(products),
		"data":   products,
	})
}

func getProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số"})
		return
	}

	for _, p := range products {
		if p.ID == id {
			c.JSON(http.StatusOK, gin.H{"status": "success", "data": p})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
}

func createProduct(c *gin.Context) {
	var newProduct Product
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu JSON không hợp lệ"})
		return
	}

	newProduct.ID = nextID
	nextID++
	products = append(products, newProduct)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Tạo sản phẩm thành công",
		"data":    newProduct,
	})
}

func updateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số"})
		return
	}

	var updated Product
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dữ liệu không hợp lệ"})
		return
	}

	for i, p := range products {
		if p.ID == id {
			updated.ID = id
			products[i] = updated
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Cập nhật thành công",
				"data":    updated,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
}

func deleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID phải là số"})
		return
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "Xóa sản phẩm thành công",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Không tìm thấy sản phẩm"})
}
