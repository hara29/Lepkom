# Academic Score Service

Aplikasi sederhana berbasis Golang untuk menghitung nilai akhir mahasiswa, menentukan grade, dan mengecek status kelulusan. Proyek ini dibuat sebagai implementasi materi **Testing, Code Review & Profiling pada Aplikasi Go**.

---

## Deskripsi

Sistem akademik ini digunakan untuk menghitung nilai akhir mahasiswa berdasarkan komponen:

| Komponen | Bobot |
|-----------|--------|
| UTS | 30% |
| UAS | 40% |
| Tugas | 30% |

Selain menghitung nilai akhir, sistem juga:

- Menentukan grade mahasiswa (A–E)
- Menentukan status kelulusan
- Melakukan validasi input nilai
- Mendukung unit testing
- Mendukung code coverage
- Mendukung benchmark testing
- Mendukung CPU profiling

---

## Struktur Project

```text
ACADEMIC_SCORE_SERVICE
│
├── go.mod
├── main.go
├── service.go
├── service_test.go
├── benchmark_test.go
└── README.md
```

### Penjelasan File

| File | Fungsi |
|---------|---------|
| service.go | Implementasi logika bisnis aplikasi |
| service_test.go | Unit test menggunakan Table Driven Test |
| benchmark_test.go | Pengujian performa (benchmark) |
| main.go | Program utama untuk menjalankan aplikasi |
| go.mod | Konfigurasi module Go |
| README.md | Dokumentasi proyek |

---

## Struktur Data

Program menggunakan struktur data berikut:

```go
type Student struct {
    Name  string
    UTS   float64
    UAS   float64
    Tugas float64
}
```

### Keterangan

| Field | Tipe | Deskripsi |
|---------|---------|---------|
| Name | string | Nama mahasiswa |
| UTS | float64 | Nilai UTS |
| UAS | float64 | Nilai UAS |
| Tugas | float64 | Nilai tugas |

---

## Fitur Utama

### 1. CalculateFinalScore()

Menghitung nilai akhir mahasiswa dengan rumus:

```
Nilai Akhir =
(UTS × 30%) +
(UAS × 40%) +
(Tugas × 30%)
```

### Validasi

- Nilai tidak boleh kurang dari 0
- Nilai tidak boleh lebih dari 100

Jika nilai tidak valid maka fungsi mengembalikan error.

Contoh:

```go
score, err := student.CalculateFinalScore()
```

---

### 2. CalculateGrade()

Mengubah nilai akhir menjadi grade.

| Nilai | Grade |
|---------|---------|
| ≥ 85 | A |
| ≥ 70 | B |
| ≥ 60 | C |
| ≥ 50 | D |
| < 50 | E |

Contoh:

```go
grade, err := student.CalculateGrade()
```

---

### 3. IsPassed()

Menentukan status kelulusan mahasiswa.

Mahasiswa dinyatakan lulus apabila memperoleh grade:

- A
- B
- C

Contoh:

```go
passed, err := student.IsPassed()
```

---

## Menjalankan Program

### Menjalankan aplikasi

```bash
go run .
```

Atau

```bash
go run main.go service.go
```

### Output

```text
Nama Mahasiswa : Cindy
Nilai Akhir    : 82
Grade          : B
Status Lulus   : true
```

---

## Unit Testing

Pengujian dilakukan menggunakan package bawaan Go yaitu:

```go
import "testing"
```

Metode yang digunakan:

- Table Driven Test
- Validasi error input

### Menjalankan seluruh test

```bash
go test
```

### Menampilkan detail test

```bash
go test -v
```

Contoh output:

```text
=== RUN   TestAllLogic_Coverage
--- PASS: TestAllLogic_Coverage (0.00s)

PASS
```

---

## Code Coverage

Code Coverage digunakan untuk mengukur persentase kode yang telah diuji.

### Menjalankan coverage

```bash
go test -v -cover service.go service_test.go
```

Contoh output:

```text
PASS
coverage: 100.0% of statements
```

### Membuat laporan coverage

```bash
go test -coverprofile=cover.out
```

File yang dihasilkan:

```text
cover.out
```

---

## Benchmark Testing

Benchmark digunakan untuk mengukur performa fungsi.

Fungsi benchmark:

```go
func BenchmarkCalculateFinalScore(b *testing.B)
```

### Menjalankan benchmark

```bash
go test -bench=.
```

Contoh output:

```text
BenchmarkCalculateFinalScore-8
100000000
3.12 ns/op
```

### Keterangan

| Parameter | Arti |
|------------|------|
| BenchmarkCalculateFinalScore-8 | Nama benchmark |
| 100000000 | Jumlah iterasi |
| 3.12 ns/op | Waktu rata-rata per operasi |

Semakin kecil nilai ns/op maka semakin baik performa fungsi.

---

## CPU Profiling

CPU Profiling digunakan untuk mengetahui bagian kode yang paling banyak menggunakan CPU.

### Menjalankan profiling

```bash
go test -v "-bench=BenchmarkCalculateFinalScore" "-cpuprofile=cpu.out" "-run=^$"
```

File yang dihasilkan:

```text
cpu.out
```

### Analisis Profiling

```bash
go tool pprof cpu.out
```

Perintah ini digunakan untuk:

- Mengidentifikasi bottleneck
- Menganalisis penggunaan CPU
- Mengoptimalkan performa aplikasi

---

## Code Review Checklist

Checklist kualitas kode yang digunakan pada proyek ini:

### Penamaan

- [x] Nama fungsi jelas
- [x] Nama variabel mudah dipahami

### Validasi Input

- [x] Nilai harus berada pada rentang 0–100

### Error Handling

- [x] Menggunakan return error
- [x] Tidak menggunakan panic

Contoh:

```go
if s.UTS < 0 || s.UTS > 100 {
    return 0, errors.New("nilai UTS tidak valid")
}
```

### Unit Test

- [x] Unit test tersedia
- [x] Menggunakan Table Driven Test

### Coverage

- [x] Coverage ≥ 70%

---

## Contoh Penggunaan

```go
student := Student{
    Name:  "Cindy",
    UTS:   80,
    UAS:   85,
    Tugas: 80,
}

score, _ := student.CalculateFinalScore()
grade, _ := student.CalculateGrade()
passed, _ := student.IsPassed()

fmt.Println(score)
fmt.Println(grade)
fmt.Println(passed)
```

Output:

```text
82
B
true
```

---

## Hasil Quality Assurance

| Kriteria | Status |
|-----------|---------|
| Unit Test PASS | ✅ |
| Coverage ≥ 70% | ✅ |
| Benchmark Berjalan | ✅ |
| CPU Profiling Berjalan | ✅ |
| Tidak Menggunakan Panic | ✅ |
| Error Handling Tersedia | ✅ |

---

## Teknologi

- Golang
- Package testing
- Benchmark Testing
- CPU Profiling
- Table Driven Test

---

## Author

Nama: Cindy Maharani

Mata Kuliah: Pemrograman Web dengan Go

Pertemuan 6 – Testing, Code Review & Profiling pada Aplikasi Go