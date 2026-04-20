//Quản lý danh sách sinh viên bằng Slice + Struct. Viết mã giả: Định nghĩa 
// struct SinhVien {Ten, Lop string; Diem float64}. Tạo slice chứa danh sách sinh viên. Viết hàm thêm, xóa, tìm kiếm theo tên.

/* Bắt đầu

Định nghĩa struct (Kiểu dữ liệu) SinhVien {
	Ten string chuỗi
	Lop string chuỗi
	Diem float64 số thực
}
 Tạo danh sách sinh viên dạng slice rỗng.
Hàm themSinhVien(danhSach, sinhVienMoi):
    thêm sinhVienMoi vào cuối danhSach
    trả về danhSach
}

Hàm xoaSinhVienTheoTen(danhSach, tenCanXoa):
    duyệt từng phần tử trong danhSach
        nếu tên sinh viên trùng với tenCanXoa
            xóa phần tử đó khỏi danhSach
            thoát vòng lặp
    trả về danhSach
}

Hàm timSinhVienTheoTen(danhSach, tenCanTim):
    duyệt từng phần tử trong danhSach
        nếu tên sinh viên trùng với tenCanTim
            trả về sinh viên tìm được và giá trị đúng
    trả về giá trị rỗng và sai
}

Trong hàm main() {
	Khởi tạo danhSachSinhVien
	Thêm một số sinh viên vào danhSachSinhVien bằng cách gọi themSinhVien()
	In ra danh sách sinh viên

	Tìm sinh viên theo tên bằng cách gọi timKiemSinhVien()
	In ra thông tin sinh viên tìm được hoặc thông báo không tìm thấy

	Xóa sinh viên theo tên bằng cách gọi xoaSinhVien()
	In ra danh sách sinh viên sau khi xóa
}  
	*/

package main

import "fmt"

// Struct SinhVien
type SinhVien struct {
	Ten  string
	Lop  string
	Diem float64
}

// Hàm thêm sinh viên vào slice
func themSinhVien(danhSach []SinhVien, sv SinhVien) []SinhVien {
	danhSach = append(danhSach, sv)
	return danhSach
}

// Hàm xóa sinh viên theo tên (xóa lần xuất hiện đầu tiên)
func xoaSinhVienTheoTen(danhSach []SinhVien, tenCanXoa string) []SinhVien {
	for i, sv := range danhSach {
		if sv.Ten == tenCanXoa {
			danhSach = append(danhSach[:i], danhSach[i+1:]...)
			break
		}
	}
	return danhSach
}

// Hàm tìm sinh viên theo tên
func timSinhVienTheoTen(danhSach []SinhVien, tenCanTim string) (SinhVien, bool) {
	for _, sv := range danhSach {
		if sv.Ten == tenCanTim {
			return sv, true
		}
	}
	return SinhVien{}, false
}

// Hàm in danh sách sinh viên
func inDanhSach(danhSach []SinhVien) {
	if len(danhSach) == 0 {
		fmt.Println("Danh sách sinh viên đang rỗng.")
		return
	}

	fmt.Println("Danh sách sinh viên:")
	for i, sv := range danhSach {
		fmt.Printf("%d. Tên: %s | Lớp: %s | Điểm: %.2f\n", i+1, sv.Ten, sv.Lop, sv.Diem)
	}
	fmt.Println()
}

func main() {
	// Tạo slice chứa danh sách sinh viên
	var danhSach []SinhVien

	// Thêm sinh viên
	danhSach = themSinhVien(danhSach, SinhVien{"An", "CNTT1", 8.5})
	danhSach = themSinhVien(danhSach, SinhVien{"Bình", "CNTT2", 7.8})
	danhSach = themSinhVien(danhSach, SinhVien{"Cường", "CNTT1", 9.1})

	// In danh sách ban đầu
	fmt.Println("=== Danh sách ban đầu ===")
	inDanhSach(danhSach)

	// Tìm sinh viên theo tên
	fmt.Println("=== Tìm sinh viên tên Bình ===")
	sv, timThay := timSinhVienTheoTen(danhSach, "Bình")
	if timThay {
		fmt.Printf("Tìm thấy: Tên: %s | Lớp: %s | Điểm: %.2f\n\n", sv.Ten, sv.Lop, sv.Diem)
	} else {
		fmt.Println("Không tìm thấy sinh viên.\n")
	}

	// Xóa sinh viên theo tên
	fmt.Println("=== Xóa sinh viên tên An ===")
	danhSach = xoaSinhVienTheoTen(danhSach, "An")

	// In lại danh sách sau khi xóa
	fmt.Println("=== Danh sách sau khi xóa ===")
	inDanhSach(danhSach)
}