package handlers

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
)

// RedirectError mengarahkan user ke halaman utama dengan pesan error pada query string.
func RedirectError(w http.ResponseWriter, r *http.Request, err error) {
	http.Redirect(w, r, fmt.Sprintf("/?error=%s", err.Error()), http.StatusSeeOther)
}

// ParsePrice mengambil nilai price dari form dan mengubahnya menjadi float64.
func ParsePrice(r *http.Request) (float64, error) {
	priceStr := r.FormValue("price")

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0, errors.New("invalid price")
	}

	return price, nil
}

// ParseStock mengambil nilai stock dari form dan mengubahnya menjadi integer.
func ParseStock(r *http.Request) (int, error) {
	stockStr := r.FormValue("stock")

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return 0, errors.New("invalid stock")
	}

	return stock, nil
}

// ReadAndValidateImage membaca file gambar dan memvalidasi ukuran serta tipe filenya.
func ReadAndValidateImage(file multipart.File, header *multipart.FileHeader) ([]byte, error) {
	defer file.Close()

	// Baca seluruh isi file gambar menjadi byte.
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, errors.New("failed to read image")
	}

	// Pastikan file gambar tidak kosong.
	if len(fileBytes) == 0 {
		return nil, errors.New("image is required")
	}

	// Batasi ukuran gambar maksimal 1MB.
	if header.Size > 1<<20 {
		return nil, errors.New("image size must be less than 1MB")
	}

	// Deteksi tipe file berdasarkan isi file, bukan hanya ekstensi.
	contentType := http.DetectContentType(fileBytes)

	if contentType != "image/jpeg" && contentType != "image/png" {
		return nil, errors.New("image must be JPEG or PNG")
	}

	return fileBytes, nil
}
