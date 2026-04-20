Phần 3: Complex Data Types & Pointers

Bài 9 — Quản lý danh sách sinh viên bằng Slice + Struct
Viết mã giả: Định nghĩa struct SinhVien {Ten, Lop string; Diem float64}. Tạo slice chứa danh sách sinh viên. Viết hàm thêm, xóa, tìm kiếm theo tên.

Bài 10 — Đếm tần suất từ bằng Map
Vẽ flowchart: Nhập 1 chuỗi văn bản, tách từ, dùng map[string]int đếm số lần xuất hiện mỗi từ. In kết quả.



Bài 11 — Swap 2 giá trị bằng Pointer
Viết mã giả hàm swap(a *int, b *int) hoán đổi giá trị 2 biến thông qua con trỏ. Vẽ sơ đồ bộ nhớ trước và sau khi swap.


Bài 12 — Cập nhật struct qua Pointer
Viết mã giả hàm tangDiem(sv *SinhVien, diemCong float64) — thay đổi điểm sinh viên thông qua pointer. Giải thích tại sao cần dùng pointer thay vì truyền giá trị.


Phần 4: Error Handling

Bài 13 — Validate tuổi người dùng
Vẽ flowchart: Hàm validateTuoi(tuoi int) error — trả về error nếu tuổi < 0 hoặc > 150, trả về nil nếu hợp lệ. Caller kiểm tra error trước khi tiếp tục.


Bài 14 — Custom Error
Viết mã giả: Tạo custom error type ValidationError {Field, Message string}. Hàm validateSinhVien(sv SinhVien) trả về ValidationError cụ thể cho từng trường không hợp lệ (tên rỗng, điểm < 0, điểm > 10).


Phần 5: Gin Framework — REST API

Bài 15 — CRUD sản phẩm (mã giả + flowchart)
Vẽ flowchart cho toàn bộ luồng API quản lý sản phẩm:

GET /products — lấy danh sách
GET /products/:id — lấy 1 sản phẩm
POST /products — tạo mới
PUT /products/:id — cập nhật
DELETE /products/:id — xóa

Dữ liệu lưu trong slice (chưa cần database). Viết mã giả cho mỗi handler.


Bài 16 — Validate request body
Vẽ flowchart: Khi POST /products, validate các trường (tên không rỗng, giá > 0, số lượng >= 0). Nếu lỗi trả về 400 + message cụ thể. Nếu hợp lệ trả về 201 + product.


Bài 17 — Tách Router và Controller
Viết mã giả mô tả cấu trúc project:
/routes      → định nghĩa route, gán handler
/controllers → xử lý logic từng endpoint
/models      → định nghĩa struct
Mô tả luồng request đi từ Router → Controller → Response.


Bài 18 — Middleware Logger + Error Handler
Vẽ flowchart: Request đến → Middleware Logger (ghi method, path, thời gian) → Handler → Nếu có error → Middleware Error Handler (trả JSON error chuẩn) → Response.
