package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 1. Middleware Logger: Ghi lại thông tin request và thời gian xử lý
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Thông tin trước khi xử lý
		fmt.Printf("--- BẮT ĐẦU: %s %s ---\n", c.Request.Method, c.Request.URL.Path)

		c.Next() // Cho phép request đi tiếp đến Handler

		// Thông tin sau khi xử lý xong
		latency := time.Since(t)
		status := c.Writer.Status()
		fmt.Printf("--- KẾT THÚC: %s %s | Status: %d | Time: %v ---\n",
			c.Request.Method, c.Request.URL.Path, status, latency)
	}
}

// 2. Middleware Error Handler: Gom lỗi và trả về JSON chuẩn
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Đợi Handler xử lý xong

		// Kiểm tra nếu có lỗi được đẩy vào c.Error()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		}
	}
}

func main() {
	r := gin.New() // Dùng gin.New() để không lấy Logger mặc định của Gin

	// Sử dụng các Middleware
	r.Use(LoggerMiddleware())
	r.Use(ErrorHandlerMiddleware())

	// Route mẫu
	r.GET("/test-success", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Thành công!"})
	})

	r.GET("/test-error", func(c *gin.Context) {
		// Thay vì trả về JSON lỗi tại đây, ta đẩy lỗi vào Context
		c.Error(errors.New("Đây là một lỗi hệ thống mô phỏng"))
	})

	r.Run(":8080")
}
