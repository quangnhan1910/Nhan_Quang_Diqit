package main

import (
	"errors"
	"fmt"
)

// Hàm validate tuổi
func validateTuoi(tuoi int) error {
	if tuoi < 0 || tuoi > 150 {
		return errors.New("Tuổi phải từ 0 đến 150")
	}
	return nil // Hợp lệ
}

func main() {
	// Caller kiểm tra error
	age := 25
	if err := validateTuoi(age); err != nil {
		fmt.Println("Lỗi:", err)
		return
	}
	fmt.Println("Tuổi hợp lệ:", age)

	// Test trường hợp lỗi
	age2 := 200
	if err := validateTuoi(age2); err != nil {
		fmt.Println("Lỗi:", err)
	}
}
