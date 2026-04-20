package routes

import (
	"my-project/controllers" // Import controllers để gán hàm xử lý

	"github.com/gin-gonic/gin"
)

// Hàm setup routes riêng cho đối tượng Product
func SetupProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		// Gắn đường dẫn "/" (thành /products) với hàm GetAllProducts trong Controller
		productGroup.GET("/", controllers.GetAllProducts)
	}
}
