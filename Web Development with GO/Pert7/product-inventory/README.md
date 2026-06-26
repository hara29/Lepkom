# Product Inventory

Project Pertemuan 7 untuk studi kasus toko Lepkom Berkah Maju. Aplikasi ini digunakan untuk mengelola data produk, menyediakan API untuk mitra, menerapkan authentication dan authorization berbasis token, serta menyediakan unit test sederhana untuk validasi endpoint yang dilindungi middleware.

## Fitur

- Login user.
- CRUD produk melalui halaman web.
- CRUD produk melalui API.
- Generate API token untuk user.
- Role authorization:
  - `admin`: dapat mengelola produk, membuat user, melihat user, dan generate token.
  - `user`: dapat mengakses/manipulasi produk menggunakan token yang valid.
- Pencatatan transaksi stok masuk/keluar melalui API.
- Middleware logger.
- Unit test untuk endpoint produk tanpa token dan dengan token.

## Struktur Project

```text
product-inventory/
├── configs/
│   └── db.go
├── handlers/
│   ├── auth_handlers.go
│   ├── product_handler.go
│   ├── serve_static.go
│   └── transaction_handler.go
├── middlewares/
│   ├── auth.go
│   └── logger.go
├── models/
│   ├── product.go
│   ├── transaction.go
│   └── user.go
├── tests/
│   └── product_test.go
├── utils/
│   ├── jwt.go
│   └── utils.go
├── .env
├── go.mod
├── go.sum
├── main.go
├── README.md
└── schema.sql
```

## Langkah Pengerjaan Sesuai ACT Pert 7

1. Menyiapkan struktur direktori project `product-inventory`.
2. Melengkapi `configs/db.go` untuk koneksi MySQL menggunakan konfigurasi dari `.env`.
3. Membuat database `inventory_db` serta tabel:
   - `users`
   - `products`
   - `stock_transactions`
4. Melengkapi `handlers/product_handler.go` untuk mengambil, menambah, mengubah, menghapus, dan mengambil detail produk berdasarkan ID.
5. Melengkapi `middlewares/auth.go` untuk validasi token dari header `Authorization`.
6. Melengkapi `middlewares/logger.go` untuk mencatat aktivitas request.
7. Membuat struct model di folder `models`, yaitu:
   - `User`
   - `Product`
   - `StockTransaction`
8. Melengkapi `tests/product_test.go` agar pengujian endpoint produk menghasilkan `PASS`.
9. Menyesuaikan file `.env` dengan konfigurasi database lokal.
10. Menguji CRUD produk melalui Website dan API.
11. Menguji perbedaan role `admin` dan `user` melalui API.

## Kebutuhan

- Go 1.22 atau lebih baru.
- MySQL/MariaDB.
- Browser.
- Aplikasi API client seperti Thunder Client, Postman, atau curl.

## Konfigurasi Database

Sesuaikan file `.env` dengan database lokal.

Contoh konfigurasi MAMP:

```env
DB_HOST=127.0.0.1
DB_PORT=8889
DB_USER=root
DB_PASSWORD=root
DB_NAME=inventory_db
JWT_SECRET_KEY=lepkom-secret-key
PORT=8087
```

Contoh konfigurasi XAMPP:

```env
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASSWORD=
DB_NAME=inventory_db
JWT_SECRET_KEY=lepkom-secret-key
PORT=8080
```

## Membuat Database dan Tabel

Jalankan file `schema.sql` di MySQL/phpMyAdmin.

Isi file tersebut akan membuat:

- database `inventory_db`
- tabel `users`
- tabel `products`
- tabel `stock_transactions`
- data awal user dan produk

Akun default:

```text
Admin:
username: admin
password: password123

User:
username: user1
password: password123
```

## Menjalankan Project

Install library/dependency yang dibutuhkan:

```bash
go get github.com/go-sql-driver/mysql@v1.8.1
go get github.com/joho/godotenv@v1.5.1
```

Jalankan aplikasi:

```bash
go run main.go
```

Buka halaman web:

```text
http://localhost:8087/login
```

Jika port di `.env` diganti menjadi `8080`, gunakan:

```text
http://localhost:8080/login
```

## Halaman Web

- `/login`: halaman login.
- `/products`: halaman CRUD produk.
- `/users`: halaman User Management dan generate API token.

## Endpoint API

### Login

```http
POST /api/login
```

Body:

```json
{
  "username": "admin",
  "password": "password123"
}
```

### Daftar Produk

```http
GET /api/products
```

Header:

```text
Authorization: Bearer <api_token>
```

### Detail Produk

```http
GET /api/products/{id}
```

### Tambah Produk

```http
POST /api/products
```

Body:

```json
{
  "name": "Pulpen",
  "stock": 25,
  "price": 3500
}
```

### Edit Produk

```http
PUT /api/products/{id}
```

Body:

```json
{
  "name": "Pulpen Biru",
  "stock": 30,
  "price": 4000
}
```

### Hapus Produk

```http
DELETE /api/products/{id}
```

### Daftar User

```http
GET /api/users
```

Endpoint ini hanya dapat diakses oleh role `admin`.

### Register User

```http
POST /api/register
```

Endpoint ini hanya dapat diakses oleh role `admin`.

Body:

```json
{
  "username": "user2",
  "password": "password123",
  "role": "user"
}
```

### Generate API Token

```http
POST /api/generate-token
```

Endpoint ini hanya dapat diakses oleh role `admin`.

Body:

```json
{
  "username": "admin"
}
```

### Transaksi Stok

```http
GET /api/transactions
POST /api/transactions
```

Body untuk transaksi:

```json
{
  "product_id": 1,
  "transaction_type": "in",
  "quantity": 10
}
```

Nilai `transaction_type` dapat berupa:

- `in`
- `out`

## Testing

Jalankan unit test:

```bash
go test ./...
```

Test yang tersedia:

- Request `GET /api/products` tanpa token harus menghasilkan `401 Unauthorized`.
- Request `GET /api/products` dengan token valid harus menghasilkan `200 OK`.

## Catatan

- Password pada project ini masih plain text karena ACT Pert 7 tidak meminta hashing password secara eksplisit.
- Untuk pengujian API dari Thunder Client/Postman, gunakan API token dari halaman `/users`.
- Jika muncul error `address already in use`, ganti nilai `PORT` di `.env` atau hentikan proses server lama yang memakai port tersebut.
