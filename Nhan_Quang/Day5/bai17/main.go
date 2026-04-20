package main

import (
	"my-project/routes" // Import routes

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Gọi hàm khởi tạo Routes
	routes.SetupProductRoutes(r)

	// Chạy server
	r.Run(":8080")
}