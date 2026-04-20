package main

import "fmt"

// Định nghĩa struct SinhVien
type SinhVien struct {
	Ten  string
	Lop  string
	Diem float64
}

// Hàm tăng điểm sử dụng Pointer
func tangDiem(sv *SinhVien, diemCong float64) {
	// Cách viết phổ biến trong Go (không cần dấu *)
	sv.Diem += diemCong

	fmt.Printf("Đã cộng %.1f điểm cho sinh viên %s\n", diemCong, sv.Ten)
	fmt.Printf("Điểm mới: %.2f\n", sv.Diem)
}

func main() {
	// Tạo sinh viên
	sv1 := SinhVien{
		Ten:  "Nguyễn Văn An",
		Lop:  "10A1",
		Diem: 7.5,
	}

	fmt.Println("=== TRƯỚC KHI CẬP NHẬT ===")
	fmt.Printf("Tên: %s\n", sv1.Ten)
	fmt.Printf("Lớp: %s\n", sv1.Lop)
	fmt.Printf("Điểm: %.2f\n\n", sv1.Diem)

	// Gọi hàm truyền pointer
	tangDiem(&sv1, 1.5)

	fmt.Println("\n=== SAU KHI CẬP NHẬT ===")
	fmt.Printf("Tên: %s\n", sv1.Ten)
	fmt.Printf("Lớp: %s\n", sv1.Lop)
	fmt.Printf("Điểm: %.2f\n", sv1.Diem)
}

// **Giải thích:**

// Trong Go, khi truyền struct vào hàm theo cách thông thường (pass by value), hàm chỉ nhận được **bản sao** của struct. Do đó, mọi thay đổi bên trong hàm không ảnh hưởng đến biến gốc bên ngoài.

// Để có thể thay đổi trực tiếp dữ liệu của struct, ta phải truyền **pointer** (`*SinhVien`).

// - Pointer chứa **địa chỉ bộ nhớ** của struct gốc.
// - Khi dùng `sv.Diem += diemCong`, hàm sẽ sửa trực tiếp trên vùng nhớ của biến gốc.
// - Kết quả là sau khi hàm chạy xong, giá trị struct bên ngoài đã thực sự được cập nhật.

// =>
// - Truyền giá trị (value): chỉ sửa bản sao → không ảnh hưởng đến biến gốc.
// - Truyền pointer: sửa trực tiếp trên dữ liệu gốc → thay đổi có hiệu lực.

// Do đó, hàm `tangDiem(sv *SinhVien, diemCong float64)` bắt buộc phải dùng pointer để cập nhật điểm sinh viên.

// ---

// Bạn có thể copy nguyên đoạn này.
// Nếu muốn ngắn hơn nữa thì nói mình chỉnh tiếp nhé!
