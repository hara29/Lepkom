package handlers

import (
	"database/sql"
	"net/http"
	"text/template"

	"toko_lepkom_npm/db"
	"toko_lepkom_npm/models"
)

// HomeView menampilkan halaman utama berisi daftar produk dari database.
func HomeView(w http.ResponseWriter, r *http.Request) {
	// Parse template layout dan index untuk halaman utama.
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/index.html",
	))
	errorMsg := r.URL.Query().Get("error")

	// Ambil data produk beserta detailnya dari tabel products dan product_details.
	rows, err := db.DB.Query(`
		SELECT 
			p.id, 
			p.name, 
			p.price,
			pd.stock, 	
			pd.is_active,
			pd.created_at
		FROM products p
		JOIN product_details pd
		ON p.id = pd.product_id
	`)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	defer rows.Close()

	var products []ProductView

	// Ubah setiap row hasil query menjadi data yang siap ditampilkan pada template.
	for rows.Next() {
		var product models.Product

		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Price,
			&product.Stock,
			&product.IsActive,
			&product.CreatedAt,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		products = append(products, ProductView{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			IsActive:  product.IsActive,
			CreatedAt: product.CreatedAt.Format("02 Jan 2006 15:04"),
		})
	}

	data := HomePageData{
		Products: products,
		Error:    errorMsg,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}

// ImageView menampilkan gambar produk berdasarkan id produk.
func ImageView(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	var image []byte

	// Ambil data gambar dari database sesuai product_id.
	err := db.DB.QueryRow("SELECT image FROM product_details WHERE product_id = ?", id).Scan(&image)

	if err != nil {
		http.NotFound(w, r)
		return
	}
	contentType := http.DetectContentType(image)
	w.Header().Set("Content-Type", contentType)
	w.Write(image)
}

// EditView menampilkan halaman form edit untuk produk yang dipilih.
func EditView(w http.ResponseWriter, r *http.Request) {
	// Parse template layout dan edit untuk halaman ubah data.
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/edit.html",
	))

	// Ambil id produk dari query string dan validasi agar tidak kosong.
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Redirect(w, r, "/?error=Id is required", http.StatusSeeOther)
		return
	}

	query := `
	SELECT 
	p.id, 
	p.name, 
	p.price,
	pd.stock, 
	pd.is_active, 
	pd.created_at
	FROM products p
	JOIN product_details pd
	ON p.id = pd.product_id
	WHERE p.id = ?
	`

	var product models.Product
	var errorMsg string

	// Ambil satu data produk yang akan diedit berdasarkan id.
	err := db.DB.QueryRow(query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.Stock,
		&product.IsActive,
		&product.CreatedAt,
	)

	if err != nil {

		// Apabila error merupakan bagian dari ErrNoRows
		if err == sql.ErrNoRows {
			errorMsg = "Data not found"
		} else {
			// Selain itu, isi dengan Internal Server Error
			errorMsg = "Internal Server Error"
		}
	}

	formatedDate := product.CreatedAt.Format("02 Jan 2006 15:04")

	data := EditPageData{
		Product: ProductView{
			Id:        product.Id,
			Name:      product.Name,
			Price:     product.Price,
			Stock:     product.Stock,
			IsActive:  product.IsActive,
			CreatedAt: formatedDate,
		},
		Error: errorMsg,
	}

	tmpl.ExecuteTemplate(w, "layout", data)
}
