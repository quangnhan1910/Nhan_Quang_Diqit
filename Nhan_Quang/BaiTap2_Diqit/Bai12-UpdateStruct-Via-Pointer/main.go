// Cập nhật struct qua Pointer. Viết mã giả hàm tangDiem(sv *SinhVien, diemCong float64) — thay đổi điểm sinh viên thông qua pointer.
// Giải thích tại sao cần dùng pointer thay vì truyền giá trị.

/*
Cấu trúc sinh viên
tên: string
lớp : string
điểm: float64
kết thúc cấu trúc sinh viên

Hàm tăng điểm (sv *sinh viên, điểm cộng float64)
	sv. điểm = sv.điểm + điểm cộng
kết thúc hàm

Hàm main()
	khởi tạo sinh viên sv với tên "Nhân", lớp "D8", điểm 8.5
	in thông tin sinh viên trước khi tăng điểm
	gọi hàm tangDiem(&sv, 1.0) để tăng điểm cho sinh viên
	in thông tin sinh viên sau khi tăng điểm
kết thúc hàm main

	Dùng pointer vì cập nhật trực tiếp từ dữ liệu gốc.
	Nếu truyền giá trị, thì hàm chỉ nhận bản sao của struct, nên khi thay đổi trong hàm,
thì hàm chỉ sửa bản sao, sẽ không làm thay đổi dữ liệu gốc.
*/

package main

import "fmt"

type SinhVien struct {
	Ten  string
	Lop  string
	Diem float64
}

func tangDiem(sv *SinhVien, diemCong float64) {
	sv.Diem = sv.Diem + diemCong
}

func main() {
	sv := SinhVien{
		Ten:  "Nhân",
		Lop:  "D8",
		Diem: 8.5,
	}

	fmt.Println("Trước khi tăng điểm:")
	fmt.Println("Tên:", sv.Ten)
	fmt.Println("Lớp:", sv.Lop)
	fmt.Println("Điểm:", sv.Diem)
	fmt.Println()

	tangDiem(&sv, 1.0)

	fmt.Println("Sau khi tăng điểm:")
	fmt.Println("Tên:", sv.Ten)
	fmt.Println("Lớp:", sv.Lop)
	fmt.Println("Điểm:", sv.Diem)
}
