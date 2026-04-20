package controllers

import (
	"my-project/models" // Import package models để dùng struct
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hàm xử lý logic cho endpoint lấy danh sách sản phẩm
func GetAllProducts(c *gin.Context) {
	// Giả lập lấy dữ liệu (thực tế sẽ lấy từ Database)
	products := []models.Product{
		{ID: 1, Name: "iPhone 15", Price: 1000},
		{ID: 2, Name: "MacBook Pro", Price: 2000},
	}

	// Đóng gói và trả về JSON
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}
