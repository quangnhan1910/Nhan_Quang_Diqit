// Hàm validateTuoi(tuoi int) error — trả về error nếu tuổi < 0 hoặc > 150,
// trả về nil nếu hợp lệ. Caller kiểm tra error trước khi tiếp tục.

package main

import (
	"errors"
	"fmt"
)

func validateTuoi(tuoi int) error {
	if tuoi < 0 || tuoi > 150 {
		return errors.New("tuoi khong hop le")
	}
	return nil
}

func main() {
	var tuoi int

	fmt.Print("Nhap tuoi: ")
	fmt.Scan(&tuoi)

	err := validateTuoi(tuoi)

	if err != nil {
		fmt.Println("Loi:", err)
		return
	}

	fmt.Println("Tuoi hop le, tiep tuc xu ly")
}
