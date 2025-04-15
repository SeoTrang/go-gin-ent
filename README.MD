# 🚀 Hướng dẫn chạy project Go (Gin + Ent)

## 📦 Yêu cầu

- Go >= 1.18
- Git
- (Tuỳ chọn) SQLite3 hoặc MySQL/Postgres nếu project dùng database khác

## 📥 Bước 1: Clone project

```bash
git clone https://github.com/SeoTrang/go-gin-ent.git
cd go-gin-ent
```

## 📄 Bước 2: Khai báo Go module (nếu chưa có)

```bash
go mod init <tên-module>
```

*Bỏ qua bước này nếu project đã có file `go.mod`.*

## 📦 Bước 3: Tải dependency

```bash
go mod tidy
```

Lệnh này sẽ tự động:
- Tải về các package còn thiếu trong `go.mod`
- Loại bỏ package không dùng nữa

## ▶️ Bước 4: Chạy project

```bash
go run main.go
```

*Nếu có nhiều file trong project, đảm bảo `main.go` hoặc file chính để khởi chạy có trong thư mục hiện tại.*

---

## ❓ Lưu ý & Câu hỏi thường gặp

### 1. Tại sao dùng `go run main.go` mà không dùng `go run`?
- Lệnh `go run` yêu cầu chỉ rõ file chính để biên dịch. Nếu chưa chỉnh rõ và thư mục chưa đúng, sẽ báo lỗi `no go files listed`.

### 2. `go mod tidy` làm gì?
- Tự động tên dự án, tải dependency, cleanup module không dùng.

### 3. Có thể build file binary không?
```bash
go build -o app_name
./app_name
```

---

## 🛠 Cấu trúc thư mục mẫu

```bash
.
├── main.go
├── go.mod
├── go.sum
├── ent/                # Code sinh ra bởi Ent
├── models/             # Các struct models
├── controllers/        # Các handler (Gin)
├── routes/             # Định tuyến API
└── README.md
```

---

Có thể bổ sung file `.env` hoặc config khác tuùy thuộc vào cách project cấu hình.

