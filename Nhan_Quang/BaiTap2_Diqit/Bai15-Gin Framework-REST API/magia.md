// Handler GET /products                        Handler = hàm xử lý request của route đó
Handler getProducts
    Nếu danh sách sản phẩm rỗng
        Trả về danh sách rỗng
    Ngược lại
        Trả về toàn bộ danh sách sản phẩm
Kết thúc handler


// Handler GET /products/:id
Handler getProductByID
    Lấy id từ URL
    Tìm sản phẩm theo id

    Nếu tìm thấy
        Trả về sản phẩm
    Ngược lại
        Trả về lỗi "Không tìm thấy sản phẩm"
Kết thúc handler


// Handler POST /products

Handler createProduct
    Đọc dữ liệu sản phẩm từ request body
    Kiểm tra dữ liệu có hợp lệ không

    Nếu dữ liệu không hợp lệ
        Trả về lỗi validate
    Ngược lại
        Tạo sản phẩm mới
        Thêm sản phẩm vào danh sách
        Trả về sản phẩm vừa tạo
Kết thúc handler


// PUT /products/:id
Handler updateProduct
    Lấy id từ URL
    Đọc dữ liệu mới từ request body
    Tìm sản phẩm theo id

    Nếu không tìm thấy
        Trả về lỗi "Không tìm thấy sản phẩm"
    Ngược lại
        Kiểm tra dữ liệu mới có hợp lệ không

        Nếu dữ liệu không hợp lệ
            Trả về lỗi validate
        Ngược lại
            Cập nhật thông tin sản phẩm
            Trả về sản phẩm sau cập nhật
Kết thúc handler


// Handler DELETE /products/:id
Handler deleteProduct
    Lấy id từ URL
    Tìm sản phẩm theo id

    Nếu không tìm thấy
        Trả về lỗi "Không tìm thấy sản phẩm"
    Ngược lại
        Xóa sản phẩm khỏi danh sách
        Trả về thông báo xóa thành công
Kết thúc handler