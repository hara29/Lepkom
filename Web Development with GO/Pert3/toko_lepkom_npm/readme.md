# CRUD Produk dengan Golang dan MySQL

Aplikasi ini merupakan implementasi CRUD (Create, Read, Update, Delete) menggunakan Golang dan MySQL.
Program digunakan untuk mengelola data produk beserta detail produk seperti stok, status aktif, dan gambar produk.

---

# Struktur Folder

```bash
a3/
│── db/
│   └── db.go
│
│── handlers/
│   ├── helper.go
│   ├── product_handler.go
│   ├── view_handler.go
│   └── view_model.go
│
│── models/
│   └── product.go
│
│── templates/
│   ├── home.html
│   └── edit.html
│
│── main.go
│── go.mod
```

---

# 1. Menjalankan XAMPP

1. Buka aplikasi XAMPP
2. Jalankan:

   * Apache
   * MySQL

Pastikan status berubah menjadi warna hijau.

---

# 2. Membuat Database di phpMyAdmin

1. Buka browser
2. Akses:

```bash
http://localhost/phpmyadmin
```

3. Klik menu **New**
4. Buat database baru dengan nama:

```sql
toko_lepkom_npm
```

5. Klik tombol **Create**

---

# 3. Membuat Table Products

Pilih database `toko_lepkom_npm`, lalu buka tab SQL dan jalankan query berikut:

```sql
CREATE TABLE products (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price DECIMAL(15,2) NOT NULL
);
```

---

# 4. Membuat Table Product Details

Masih pada tab SQL, jalankan query berikut:

```sql
CREATE TABLE product_details (
    product_id VARCHAR(100) PRIMARY KEY,
    stock INT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    image LONGBLOB,

    CONSTRAINT fk_product
    FOREIGN KEY (product_id)
    REFERENCES products(id)
    ON DELETE CASCADE
);
```

---

# 5. Download atau Clone Project

## Jika menggunakan ZIP

1. Download folder project
2. Extract project
3. Buka folder project menggunakan VSCode

---

## Jika menggunakan Git

```bash
git clone <link_repository>
```

Lalu masuk ke folder project:

```bash
cd a3
```

---

# 6. Install Dependency

Buka terminal VSCode lalu jalankan:

```bash
go mod tidy
```

atau:

```bash
go get github.com/go-sql-driver/mysql
```

---

# 7. Konfigurasi Database

Buka file:

```bash
db/db.go
```

Sesuaikan konfigurasi database:

```go
dsn := "root:@tcp(127.0.0.1:3306)/toko_lepkom_npm"
```

Keterangan:

* root → username MySQL
* kosong setelah titik dua → password MySQL
* toko_lepkom_npm → nama database

Jika menggunakan password MySQL:

```go
dsn := "root:password@tcp(127.0.0.1:3306)/toko_lepkom_npm"
```

---

# 8. Menjalankan Program

Buka terminal pada folder project lalu jalankan:

```bash
go run main.go
```

Jika berhasil akan muncul:

```bash
Server running on :8080
```

---

# 9. Membuka Aplikasi

Buka browser lalu akses:

```bash
http://localhost:8080
```

---

# 10. Fitur Aplikasi

Aplikasi memiliki fitur:

* Menampilkan semua produk
* Menambahkan produk
* Mengedit produk
* Menghapus produk
* Menampilkan gambar produk
* Validasi upload gambar
* Error handling

---

# 11. Validasi Upload Gambar

Program hanya menerima:

* JPG / JPEG
* PNG

Dengan ukuran maksimal:

```bash
1 MB
```

---

# 12. Route Program

| Route       | Fungsi             |
| ----------- | ------------------ |
| /           | Halaman home       |
| /image?id=1 | Menampilkan gambar |
| /edit?id=1  | Halaman edit       |
| /create     | Menambah produk    |
| /update     | Mengupdate produk  |
| /delete     | Menghapus produk   |

---

# 13. Screenshot Hasil Pengerjaan

Berikut hasil yang harus didokumentasikan dalam laporan:

## Database

* Database `toko_lepkom_npm`
* Table `products`
* Table `product_details`

## Program Berjalan

* Terminal berhasil menjalankan `go run main.go`
* Tampilan halaman home

## CRUD

* Menambahkan data produk
* Data berhasil tampil di tabel
* Edit produk berhasil
* Delete produk berhasil

## Upload Gambar

* Gambar berhasil tampil
* Validasi format gambar
* Validasi ukuran gambar lebih dari 1MB

## Error Handling

* Error input price
* Error input stock
* Error upload file tidak valid

---

# 14. Format Pengumpulan

1. File laporan PDF

```bash
Pert3_Nama_NPM.pdf
```

3. Screenshot dokumentasi hasil pengerjaan

---

# 15. Author

Nama  : [Isi Nama]
NPM   : [Isi NPM]
Kelas : [Isi Kelas]
