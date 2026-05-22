# Pertemuan 5 - REST API dan Integrasi Frontend Lit UI

Proyek ini merupakan hasil pengerjaan tugas **ACT Pert 5: REST API dan Integrasi Frontend (Lit UI)**. Aplikasi yang dibuat adalah **Task Manager** sederhana dengan backend REST API menggunakan Go dan frontend berbasis Lit.

Melalui aplikasi ini, pengguna dapat melihat daftar task, menambahkan task baru, mengubah task, dan menghapus task melalui tampilan web yang terhubung langsung ke REST API.

## Fitur

- REST API menggunakan package standar Go `net/http`
- CRUD task:
  - Menampilkan seluruh task
  - Menampilkan task berdasarkan ID
  - Menambahkan task
  - Mengubah seluruh data task
  - Mengubah sebagian data task
  - Menghapus task
- Validasi status task
- Response API dalam format JSON yang konsisten
- Middleware logger untuk mencatat request
- Frontend Lit UI untuk integrasi langsung dengan API
- Static file server untuk menyajikan halaman frontend

## Teknologi

- Go
- net/http
- HTML
- JavaScript
- Lit

## Struktur Folder

```text
.
├── go.mod
├── main.go
├── handlers
│   └── task_handler.go
├── middlewares
│   └── logger.go
├── models
│   └── task.go
├── utils
│   └── response.go
└── static
    ├── index.html
    └── lit-all.min.js
```

Keterangan:

- `main.go`: entry point aplikasi, routing API, dan static file server.
- `handlers/task_handler.go`: handler untuk semua operasi CRUD task.
- `models/task.go`: definisi struct `Task`.
- `utils/response.go`: helper untuk membuat response JSON.
- `middlewares/logger.go`: middleware untuk logging request.
- `static/index.html`: frontend Task Manager menggunakan Lit.
- `static/lit-all.min.js`: file Lit lokal yang disertakan dalam folder static.

## Model Data

Task memiliki struktur data berikut:

```json
{
  "id": 1,
  "title": "Belajar REST API",
  "description": "Mengerjakan tugas Pertemuan 5",
  "status": "pending",
  "created_at": "2026-05-21T10:00:00Z"
}
```

Status task yang valid:

- `pending`
- `in-progress`
- `done`

## Cara Menjalankan

Pastikan Go sudah terpasang pada komputer.

1. Masuk ke folder proyek:

```bash
cd "/Users/cindymaharani/Documents/go/src/Lepkom/Web Development Go/Pert5"
```

2. Jalankan aplikasi:

```bash
go run main.go
```

3. Buka aplikasi di browser:

```text
http://localhost:8080
```

Server akan berjalan pada:

```text
http://localhost:8080
```

## Endpoint API

Base URL:

```text
http://localhost:8080
```

| Method | Endpoint | Deskripsi |
| --- | --- | --- |
| GET | `/api/tasks` | Mengambil semua task |
| POST | `/api/tasks` | Membuat task baru |
| GET | `/api/tasks/{id}` | Mengambil detail task berdasarkan ID |
| PUT | `/api/tasks/{id}` | Mengubah seluruh data task berdasarkan ID |
| PATCH | `/api/tasks/{id}` | Mengubah sebagian data task berdasarkan ID |
| DELETE | `/api/tasks/{id}` | Menghapus task berdasarkan ID |

## Format Response API

Semua response API menggunakan format berikut:

```json
{
  "status": "success",
  "message": "Task created",
  "data": {
    "id": 1,
    "title": "Belajar Go",
    "description": "Membuat REST API sederhana",
    "status": "pending",
    "created_at": "2026-05-21T10:00:00Z"
  }
}
```

Jika terjadi error, response tetap menggunakan struktur yang sama:

```json
{
  "status": "error",
  "message": "Invalid status",
  "data": null
}
```

## Contoh Penggunaan API

### 1. Mengambil Semua Task

```bash
curl http://localhost:8080/api/tasks
```

### 2. Membuat Task Baru

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go",
    "description": "Membuat REST API sederhana",
    "status": "pending"
  }'
```

### 3. Mengambil Task Berdasarkan ID

```bash
curl http://localhost:8080/api/tasks/1
```

### 4. Mengubah Seluruh Data Task

```bash
curl -X PUT http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Belajar Go dan Lit",
    "description": "Mengintegrasikan REST API dengan frontend",
    "status": "in-progress"
  }'
```

### 5. Mengubah Sebagian Data Task

```bash
curl -X PATCH http://localhost:8080/api/tasks/1 \
  -H "Content-Type: application/json" \
  -d '{
    "status": "done"
  }'
```

### 6. Menghapus Task

```bash
curl -X DELETE http://localhost:8080/api/tasks/1
```

## Frontend

Frontend berada pada file `static/index.html` dan dibuat menggunakan Lit. Komponen utama yang digunakan adalah custom element:

```html
<task-app></task-app>
```

Frontend melakukan request ke endpoint:

- `GET /api/tasks`
- `POST /api/tasks`
- `PUT /api/tasks/{id}`
- `DELETE /api/tasks/{id}`

Melalui tampilan web, pengguna dapat:

- Melihat daftar task
- Menambahkan task baru
- Mengedit task
- Menghapus task

## Catatan Implementasi

- Data task disimpan di memory menggunakan slice Go, sehingga data akan hilang ketika server dihentikan.
- ID task dibuat otomatis menggunakan variabel `nextID`.
- Field `created_at` dibuat otomatis saat task baru berhasil ditambahkan.
- Status task divalidasi agar hanya menerima `pending`, `in-progress`, atau `done`.
- Middleware logger akan menampilkan method, path, dan durasi request pada terminal.

## Credit

Code ini dibuat oleh:

**Dhirsya Ramdhan Pattah**
