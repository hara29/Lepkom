package handlers

import (
	"net/http"

	"toko_lepkom_npm/db"
)

// CreateProductHandler menangani proses penambahan produk baru dari form.
func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	// Lakukan validasi method yang dikirim apabila bukan POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//Lakukan validasi ketika file yang dikirim lebih besar dari 1MB
	r.ParseMultipartForm(1 << 20)

	id := r.FormValue("id")
	name := r.FormValue("name")

	// Lakukan konversi tipe data menjadi float terhadap price
	price, err := ParsePrice(r)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	// Lakukan konversi tipe data menjadi int terhadap stock
	stock, err := ParseStock(r)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	var imageBytes []byte

	// Lakukan pengambilan gambar
	file, header, err := r.FormFile("image")

	// Jika user upload gambar
	if err == nil {

		// Validasi gambar
		imageBytes, err = ReadAndValidateImage(file, header)

		if err != nil {
			RedirectError(w, r, err)
			return
		}
	}
	// Masukkan data ke tabel products
	_, err = db.DB.Exec(
		"INSERT INTO products(id, name, price) VALUES(?, ?, ?)",
		id,
		name,
		price,
	)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	// Masukkan data ke tabel product_details
	_, err = db.DB.Exec(
		`INSERT INTO product_details
		(product_id, stock, image)
		VALUES(?, ?, ?)`,
		id,
		stock,
		imageBytes,
	)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// UpdateProductHandler menangani proses perubahan data produk berdasarkan id.
func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	// Lakukan pengecekan apakah method yang dikirim adalah POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Lakukan validasi apakah file yang dikirim tidal lebih dari 1MB
	r.ParseMultipartForm(1 << 20)

	id := r.FormValue("id")
	name := r.FormValue("name")

	// Lakukan konversi tipe data menjadi float terhadap price
	price, err := ParsePrice(r)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	// Lakukan konversi tipe data menjadi int terhadap stock
	stock, err := ParseStock(r)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	isActive := r.FormValue("is_active")

	var imageBytes []byte

	file, header, err := r.FormFile("image")

	// Lakukan pengecekan apabila err == nil
	if err == nil {

		// maka selanjutnya validasi gambar
		imageBytes, err = ReadAndValidateImage(file, header)

		if err != nil {
			RedirectError(w, r, err)
			return
		}
	}

	// Update tabel products
	_, err = db.DB.Exec(
		"UPDATE products SET name=?, price=? WHERE id=?",
		name,
		price,
		id,
	)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	// Update tabel product_details
	if len(imageBytes) > 0 {

		_, err = db.DB.Exec(
			`UPDATE product_details
			SET stock=?, is_active=?, image=?
			WHERE product_id=?`,
			stock,
			isActive == "true",
			imageBytes,
			id,
		)

	} else {

		_, err = db.DB.Exec(
			`UPDATE product_details
			SET stock=?, is_active=?
			WHERE product_id=?`,
			stock,
			isActive == "true",
			id,
		)
	}

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteProductHandler menangani proses penghapusan produk berdasarkan id.
func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	// Lakukan pengecekan apakah method yang dikirim adalah POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()

	id := r.FormValue("id")

	// Hapus data dari table products dengan menggunakan id
	_, err := db.DB.Exec(
		"DELETE FROM products WHERE id = ?",
		id,
	)

	if err != nil {
		RedirectError(w, r, err)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
