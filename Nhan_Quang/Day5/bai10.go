package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== ĐẾM TẦN SUẤT TỪ BẰNG MAP ===")

	// Nhập chuỗi văn bản
	fmt.Print("Nhập một chuỗi văn bản: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	// Tách từ và chuẩn hóa (chuyển về chữ thường, loại bỏ dấu câu nếu cần)
	words := strings.Fields(strings.ToLower(text))

	// Đếm tần suất bằng map
	tanSuat := make(map[string]int)
	for _, word := range words {
		// Loại bỏ dấu câu cơ bản nếu muốn (tùy chọn)
		word = strings.Trim(word, ".,!?;:'\"")
		if word != "" {
			tanSuat[word]++
		}
	}

	// In kết quả
	fmt.Println("\nKết quả đếm tần suất từ:")
	fmt.Println(strings.Repeat("-", 40))
	fmt.Printf("%-20s %s\n", "Từ", "Số lần xuất hiện")
	fmt.Println(strings.Repeat("-", 40))

	for word, count := range tanSuat {
		fmt.Printf("%-20s %d\n", word, count)
	}
}