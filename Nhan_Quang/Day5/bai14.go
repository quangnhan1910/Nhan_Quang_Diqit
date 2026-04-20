package main

import (
	"fmt"
)

// Định nghĩa struct SinhVien
type SinhVien struct {
	Ten  string
	Lop  string
	Diem float64
}

// Định nghĩa Custom Error
type ValidationError struct {
	Field   string
	Message string
}

// Implement interface error
func (e ValidationError) Error() string {
	return fmt.Sprintf("Lỗi trường [%s]: %s", e.Field, e.Message)
}

// Hàm validate sinh viên với custom error
func validateSinhVien(sv SinhVien) error {
	if sv.Ten == "" {
		return ValidationError{
			Field:   "Ten",
			Message: "Tên không được để trống",
		}
	}

	if sv.Diem < 0 {
		return ValidationError{
			Field:   "Diem",
			Message: "Điểm không được nhỏ hơn 0",
		}
	}

	if sv.Diem > 10 {
		return ValidationError{
			Field:   "Diem",
			Message: "Điểm không được lớn hơn 10",
		}
	}

	return nil // Hợp lệ
}

func main() {
	// Test các trường hợp
	tests := []SinhVien{
		{Ten: "", Lop: "10A1", Diem: 8.5},           // Lỗi tên rỗng
		{Ten: "Nguyễn Văn An", Lop: "10A1", Diem: -1}, // Lỗi điểm âm
		{Ten: "Trần Thị Bích", Lop: "10A2", Diem: 11}, // Lỗi điểm > 10
		{Ten: "Lê Hoàng Nam", Lop: "10A1", Diem: 9.0}, // Hợp lệ
	}

	for i, sv := range tests {
		fmt.Printf("=== Test %d: %s ===\n", i+1, sv.Ten)
		if err := validateSinhVien(sv); err != nil {
			fmt.Println("Lỗi:", err)
		} else {
			fmt.Println("Sinh viên hợp lệ!")
		}
		fmt.Println()
	}
}