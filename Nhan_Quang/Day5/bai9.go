
package main

import (
	"fmt"
	"strings"
)

// Định nghĩa struct SinhVien
type SinhVien struct {
	Ten  string
	Lop  string
	Diem float64
}

// Hàm thêm sinh viên vào slice
func ThemSinhVien(ds *[]SinhVien, sv SinhVien) {
	*ds = append(*ds, sv)
	fmt.Printf("Đã thêm sinh viên: %s\n", sv.Ten)
}

// Hàm xóa sinh viên theo tên
func XoaSinhVien(ds *[]SinhVien, ten string) bool {
	for i, sv := range *ds {
		if strings.EqualFold(sv.Ten, ten) { // So sánh không phân biệt hoa thường
			*ds = append((*ds)[:i], (*ds)[i+1:]...)
			fmt.Printf("Đã xóa sinh viên: %s\n", ten)
			return true
		}
	}
	fmt.Printf("Không tìm thấy sinh viên tên: %s\n", ten)
	return false
}

// Hàm tìm kiếm sinh viên theo tên (tìm gần đúng)
func TimKiemTheoTen(ds []SinhVien, ten string) []SinhVien {
	var ketQua []SinhVien
	for _, sv := range ds {
		if strings.Contains(strings.ToLower(sv.Ten), strings.ToLower(ten)) {
			ketQua = append(ketQua, sv)
		}
	}
	return ketQua
}

// Hàm in danh sách sinh viên
func InDanhSach(ds []SinhVien) {
	if len(ds) == 0 {
		fmt.Println("Danh sách sinh viên trống!")
		return
	}

	fmt.Println("\n=== DANH SÁCH SINH VIÊN ===")
	fmt.Printf("%-5s %-20s %-10s %s\n", "STT", "Họ tên", "Lớp", "Điểm")
	fmt.Println(strings.Repeat("-", 50))

	for i, sv := range ds {
		fmt.Printf("%-5d %-20s %-10s %.2f\n", i+1, sv.Ten, sv.Lop, sv.Diem)
	}
	fmt.Println(strings.Repeat("-", 50))
}

func main() {
	// Tạo slice chứa danh sách sinh viên
	var danhSach []SinhVien

	// Thêm sinh viên
	ThemSinhVien(&danhSach, SinhVien{"Nguyễn Văn An", "10A1", 8.5})
	ThemSinhVien(&danhSach, SinhVien{"Trần Thị Bích", "10A2", 9.0})
	ThemSinhVien(&danhSach, SinhVien{"Lê Hoàng Nam", "10A1", 7.5})
	ThemSinhVien(&danhSach, SinhVien{"Phạm Minh Quân", "10A3", 8.8})

	InDanhSach(danhSach)

	// Tìm kiếm sinh viên
	fmt.Println("\nTìm kiếm sinh viên tên 'An':")
	ketQua := TimKiemTheoTen(danhSach, "An")
	for _, sv := range ketQua {
		fmt.Printf("- %s - %s - %.2f\n", sv.Ten, sv.Lop, sv.Diem)
	}

	// Xóa sinh viên
	XoaSinhVien(&danhSach, "Trần Thị Bích")
	InDanhSach(danhSach)

	// Thử xóa sinh viên không tồn tại
	XoaSinhVien(&danhSach, "ABC XYZ")
}
