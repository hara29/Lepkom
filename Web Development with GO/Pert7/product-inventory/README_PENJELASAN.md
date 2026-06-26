# Penjelasan Project Product Inventory

Dokumen ini berisi penjelasan project `product-inventory` dengan bahasa yang lebih sederhana. Tujuannya agar mahasiswa tingkat awal dapat memahami alur kerja project, fungsi setiap folder, dan hubungan antara web, API, database, authentication, authorization, serta testing.

## Gambaran Besar

Project ini adalah aplikasi sederhana untuk mengelola produk toko. Aplikasi ini dapat digunakan melalui halaman web dan juga melalui API.

Bayangkan ada sebuah toko yang memiliki banyak produk. Toko tersebut membutuhkan sistem untuk:

- mencatat produk,
- melihat daftar produk,
- menambah produk baru,
- mengubah data produk,
- menghapus produk,
- membuat token API untuk mitra,
- membatasi akses berdasarkan login dan role user.

Project `product-inventory` dibuat untuk menyelesaikan kebutuhan tersebut.

## Bagian Utama Project

Secara sederhana, project ini terdiri dari 4 bagian besar:

1. Web UI
2. API
3. Database
4. Authentication dan Authorization

## 1. Web UI

Web UI adalah halaman yang dibuka melalui browser.

Contoh halaman:

```text
/login
/products
/users
```

Halaman `/login` digunakan untuk masuk ke sistem.

Halaman `/products` digunakan untuk mengelola produk, seperti melihat, menambah, mengedit, dan menghapus produk.

Halaman `/users` digunakan untuk user management dan generate API token.

## 2. API

API adalah jalur komunikasi yang digunakan oleh aplikasi lain, misalnya Thunder Client, Postman, atau aplikasi mitra toko.

Contoh endpoint API:

```text
GET    /api/products
GET    /api/products/{id}
POST   /api/products
PUT    /api/products/{id}
DELETE /api/products/{id}
```

API ini membuat data produk bisa diakses tanpa harus membuka halaman web.

## 3. Database

Database digunakan untuk menyimpan data secara permanen.

Project ini menggunakan MySQL dengan database:

```text
inventory_db
```

Tabel yang digunakan:

```text
users
products
stock_transactions
```

Tabel `users` menyimpan data user.

Tabel `products` menyimpan data produk.

Tabel `stock_transactions` menyimpan data transaksi stok masuk dan keluar.

## 4. Authentication dan Authorization

Authentication dan authorization adalah bagian keamanan aplikasi.

Authentication berarti mengecek identitas user.

Pertanyaannya:

```text
Kamu siapa?
```

Contohnya login menggunakan:

```text
username: admin
password: password123
```

Authorization berarti mengecek hak akses user.

Pertanyaannya:

```text
Kamu boleh melakukan apa?
```

Contohnya:

- admin boleh membuat user baru,
- admin boleh generate token,
- user biasa tidak boleh membuat user baru,
- user biasa tetap boleh mengakses produk jika token valid.

## Alur Kerja Aplikasi

Alur kerja project ini dapat digambarkan seperti berikut:

```text
Browser / API Client
        ↓
Route di main.go
        ↓
Middleware
        ↓
Handler
        ↓
Database
        ↓
Response
```

Penjelasannya:

1. User membuka browser atau mengirim request dari API client.
2. Request masuk ke route yang ada di `main.go`.
3. Middleware mengecek apakah request memiliki token yang valid.
4. Jika valid, request diteruskan ke handler.
5. Handler menjalankan proses, misalnya mengambil data produk dari database.
6. Server mengirim response kembali ke browser atau API client.

## Penjelasan Folder

## Folder `configs`

Folder ini berisi konfigurasi project.

File utama:

```text
configs/db.go
```

File ini digunakan untuk menghubungkan aplikasi Go dengan database MySQL.

Konfigurasi database dibaca dari file `.env`, misalnya:

```env
DB_HOST=127.0.0.1
DB_PORT=8889
DB_USER=root
DB_PASSWORD=root
DB_NAME=inventory_db
```

Data tersebut dipakai untuk membuat koneksi ke database.

## Folder `models`

Folder `models` berisi bentuk data yang digunakan dalam aplikasi.

Contoh model produk:

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Stock int     `json:"stock"`
	Price float64 `json:"price"`
}
```

Struct tersebut menggambarkan satu data produk.

Jika di database ada tabel `products`, maka di Go kita membuat struct `Product` untuk mewakili data dari tabel tersebut.

## Folder `handlers`

Folder `handlers` berisi fungsi untuk menangani request.

Contoh fungsi handler:

```text
GetProducts
CreateProduct
UpdateProduct
DeleteProduct
GetProductByID
```

Misalnya user mengirim request:

```http
GET /api/products
```

Maka handler `GetProducts` akan dijalankan untuk mengambil daftar produk dari database.

## Folder `middlewares`

Folder `middlewares` berisi fungsi perantara sebelum request masuk ke handler.

Middleware dapat dianggap sebagai penjaga.

Contohnya middleware auth:

```text
AuthMiddleware
```

Tugasnya:

- mengecek apakah request membawa token,
- mengecek apakah token valid,
- menolak request jika token tidak ada atau tidak valid.

Jika request tidak membawa token, server akan mengirim response:

```json
{
  "status": "error",
  "message": "Mana Tokennya?"
}
```

Dengan status:

```text
401 Unauthorized
```

## Folder `utils`

Folder `utils` berisi fungsi bantu yang sering digunakan.

Contohnya:

- membuat response JSON,
- membuat token,
- memvalidasi token.

Fungsi di folder ini membantu agar kode di handler dan middleware tidak terlalu panjang.

## Folder `tests`

Folder `tests` berisi file testing.

Testing digunakan untuk memastikan program berjalan sesuai aturan.

Contohnya:

```text
GET /api/products tanpa token harus ditolak
```

Maka test akan mengirim request tanpa token.

Jika server mengembalikan status:

```text
401 Unauthorized
```

maka test dianggap benar dan hasilnya `PASS`.

## Penjelasan Token

Token adalah seperti kartu akses.

Setelah user login, server memberikan token. Token tersebut disimpan di browser atau digunakan di API client.

Saat mengakses API, token dikirim melalui header:

```text
Authorization: Bearer <token>
```

Jika token valid, request akan dilanjutkan.

Jika token tidak valid, request akan ditolak.

## Contoh Alur Login

1. User membuka halaman `/login`.
2. User mengisi username dan password.
3. Browser mengirim request ke:

```http
POST /api/login
```

4. Server mengecek username dan password di database.
5. Jika benar, server membuat token.
6. Token disimpan di browser.
7. User diarahkan ke halaman produk.

## Contoh Alur Melihat Produk

1. User membuka halaman `/products`.
2. Browser mengirim request ke:

```http
GET /api/products
```

3. Token dikirim di header.
4. Middleware mengecek token.
5. Jika token valid, handler mengambil data dari tabel `products`.
6. Data produk dikirim kembali dalam format JSON.
7. Browser menampilkan data produk dalam tabel.

## Contoh Alur Tambah Produk

1. User mengisi form produk.
2. User klik tombol Simpan.
3. Browser mengirim request:

```http
POST /api/products
```

4. Middleware mengecek token.
5. Handler membaca data dari body request.
6. Data disimpan ke tabel `products`.
7. Server mengirim response bahwa produk berhasil ditambahkan.

## Contoh Alur Edit Produk

1. User klik tombol Edit pada salah satu produk.
2. Data produk masuk ke form.
3. User mengubah nama, stok, atau harga.
4. User klik tombol Simpan.
5. Browser mengirim request:

```http
PUT /api/products/{id}
```

6. Handler memperbarui data produk di database.
7. Data produk terbaru ditampilkan kembali.

## Contoh Alur Hapus Produk

1. User klik tombol Hapus.
2. Browser mengirim request:

```http
DELETE /api/products/{id}
```

3. Middleware mengecek token.
4. Handler menghapus data produk dari database.
5. Tabel produk diperbarui.

## Penjelasan Role Admin dan User

Project ini memiliki dua role:

```text
admin
user
```

Admin memiliki akses lebih besar.

Admin dapat:

- melihat produk,
- menambah produk,
- mengedit produk,
- menghapus produk,
- melihat daftar user,
- membuat user baru,
- generate API token.

User biasa dapat:

- melihat produk,
- menambah produk,
- mengedit produk,
- menghapus produk.

User biasa tidak boleh membuat user baru.

Jika user biasa mencoba membuat user baru melalui endpoint:

```http
POST /api/register
```

maka server akan menolak request tersebut.

## Kenapa Ada API Token?

API token digunakan agar aplikasi luar atau mitra toko dapat mengakses data produk.

Contoh:

Mitra ingin melihat daftar produk tanpa membuka website. Maka mitra bisa memakai endpoint:

```http
GET /api/products
```

Tetapi mitra tetap harus membawa token:

```text
Authorization: Bearer <api_token>
```

Dengan begitu, API tidak bisa diakses sembarang orang.

## Kenapa `go test` Menghasilkan PASS?

Pada ACT Pert 7, kondisi awal test dibuat gagal karena ekspektasinya belum benar.

Request tanpa token seharusnya ditolak.

Artinya response yang benar adalah:

```text
401 Unauthorized
```

Jika test mengecek bahwa hasilnya harus `401`, maka test akan `PASS`.

Jadi `PASS` bukan berarti API bisa diakses tanpa token.

Sebaliknya, `PASS` berarti middleware berhasil menolak request tanpa token.

## Analogi Sederhana

Bayangkan aplikasi ini seperti toko.

Database adalah gudang penyimpanan data.

Model adalah format kartu data barang atau user.

Handler adalah pegawai yang melayani permintaan.

Middleware auth adalah satpam yang mengecek kartu akses.

Token adalah kartu akses.

API adalah loket khusus untuk aplikasi luar.

Web UI adalah meja pelayanan untuk user manusia.

## Inti yang Perlu Dipahami

Hal terpenting dari project ini adalah memahami alur:

```text
Request masuk
→ route menentukan tujuan
→ middleware mengecek keamanan
→ handler memproses request
→ database dibaca atau diubah
→ response dikirim kembali
```

Jika alur tersebut sudah dipahami, maka struktur project akan lebih mudah dimengerti.

Setiap folder memiliki tugas masing-masing. Project ini bukan hanya tentang menulis banyak file Go, tetapi tentang menyusun bagian-bagian kecil agar menjadi satu sistem web yang utuh.
