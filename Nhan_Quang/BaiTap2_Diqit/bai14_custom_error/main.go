// Bài 14 — Custom Error.
// Viết mã giả: Tạo custom error type ValidationError {Field, Message string}.
// Hàm validateSinhVien(sv SinhVien) trả về ValidationError cụ thể cho từng trường không hợp lệ (tên rỗng, điểm < 0, điểm > 10).

/*
Cấu trúc ValidationError
    Field: string
    Message: string
Kết thúc cấu trúc

Hàm validateSinhVien(sv SinhVien)
    Nếu tên rỗng
        Trả về ValidationError(Field = "Tên", Message = "Tên không được để trống")

    Nếu điểm < 0
        Trả về ValidationError(Field = "Điểm", Message = "Điểm không được nhỏ hơn 0")

    Nếu điểm > 10
        Trả về ValidationError(Field = "Điểm", Message = "Điểm không được lớn hơn 10")

    Trả về nil
Kết thúc hàm
*/

package main

import "fmt"

type SinhVien struct {
	Ten  string
	Diem float64
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func validateSinhVien(sv SinhVien) error {
	if sv.Ten == "" {
		return ValidationError{
			Field:   "Tên",
			Message: "Tên không được để trống",
		}
	}

	if sv.Diem < 0 {
		return ValidationError{
			Field:   "Điểm",
			Message: "Điểm không được nhỏ hơn 0",
		}
	}

	if sv.Diem > 10 {
		return ValidationError{
			Field:   "Điểm",
			Message: "Điểm không được lớn hơn 10",
		}
	}

	return nil
}

func main() {
	sv := SinhVien{
		Ten:  "", // ngay đây sẽ gây lỗi vì tên rỗng
		Diem: 11, // ngay đây cũng sẽ gây lỗi vì điểm lớn hơn 10
	}

	err := validateSinhVien(sv)

	if err != nil {
		fmt.Println("Lỗi:", err)
	} else {
		fmt.Println("Sinh viên hợp lệ")
	}
}
