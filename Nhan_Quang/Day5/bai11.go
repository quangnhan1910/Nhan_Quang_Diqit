package main

import "fmt"

// Hàm swap sử dụng pointer
func swap(a *int, b *int) {
	temp := *a // Lưu giá trị của a
	*a = *b    // Gán giá trị của b vào a
	*b = temp  // Gán giá trị cũ của a vào b
}

func main() {
	x := 5
	y := 10

	fmt.Println("Trước khi swap:")
	fmt.Printf("x = %d, y = %d\n", x, y)
	fmt.Printf("Địa chỉ của x: %p, Địa chỉ của y: %p\n", &x, &y)

	// Gọi hàm swap truyền địa chỉ
	swap(&x, &y)

	fmt.Println("\nSau khi swap:")
	fmt.Printf("x = %d, y = %d\n", x, y)
}
