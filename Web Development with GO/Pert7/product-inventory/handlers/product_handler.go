package handlers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"product-inventory/configs"
	"product-inventory/models"
	"product-inventory/utils"
)

var (
	productMu     sync.Mutex
	products      = []models.Product{{ID: 1, Name: "Pensil", Stock: 30, Price: 2500}, {ID: 2, Name: "Buku Tulis", Stock: 20, Price: 5000}}
	nextProductID = 3
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if configs.DB != nil {
		rows, err := configs.DB.Query("SELECT id, name, stock, price FROM products ORDER BY id")
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal mengambil data produk", nil)
			return
		}
		defer rows.Close()

		var result []models.Product
		for rows.Next() {
			var p models.Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Stock, &p.Price); err != nil {
				utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membaca data produk", nil)
				return
			}
			result = append(result, p)
		}

		utils.WriteJSON(w, http.StatusOK, "success", "Daftar produk", result)
		return
	}

	productMu.Lock()
	defer productMu.Unlock()
	utils.WriteJSON(w, http.StatusOK, "success", "Daftar produk", products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var input models.Product
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Body request tidak valid", nil)
		return
	}
	if strings.TrimSpace(input.Name) == "" || input.Stock < 0 || input.Price < 0 {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Nama, stok, dan harga wajib valid", nil)
		return
	}

	if configs.DB != nil {
		result, err := configs.DB.Exec("INSERT INTO products (name, stock, price) VALUES (?, ?, ?)", input.Name, input.Stock, input.Price)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal menambah produk", nil)
			return
		}
		id, _ := result.LastInsertId()
		input.ID = int(id)
		utils.WriteJSON(w, http.StatusCreated, "success", "Produk berhasil ditambahkan", input)
		return
	}

	productMu.Lock()
	input.ID = nextProductID
	nextProductID++
	products = append(products, input)
	productMu.Unlock()

	utils.WriteJSON(w, http.StatusCreated, "success", "Produk berhasil ditambahkan", input)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, ok := productIDFromPath(w, r)
	if !ok {
		return
	}

	if configs.DB != nil {
		var product models.Product
		err := configs.DB.QueryRow("SELECT id, name, stock, price FROM products WHERE id = ?", id).
			Scan(&product.ID, &product.Name, &product.Stock, &product.Price)
		if err == sql.ErrNoRows {
			utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
			return
		}
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal mengambil produk", nil)
			return
		}
		utils.WriteJSON(w, http.StatusOK, "success", "Produk ditemukan", product)
		return
	}

	productMu.Lock()
	defer productMu.Unlock()
	for _, product := range products {
		if product.ID == id {
			utils.WriteJSON(w, http.StatusOK, "success", "Produk ditemukan", product)
			return
		}
	}
	utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := productIDFromPath(w, r)
	if !ok {
		return
	}

	var input models.Product
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Body request tidak valid", nil)
		return
	}
	if strings.TrimSpace(input.Name) == "" || input.Stock < 0 || input.Price < 0 {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Nama, stok, dan harga wajib valid", nil)
		return
	}
	input.ID = id

	if configs.DB != nil {
		result, err := configs.DB.Exec("UPDATE products SET name = ?, stock = ?, price = ? WHERE id = ?", input.Name, input.Stock, input.Price, id)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal mengubah produk", nil)
			return
		}
		if affected, _ := result.RowsAffected(); affected == 0 {
			utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
			return
		}
		utils.WriteJSON(w, http.StatusOK, "success", "Produk berhasil diubah", input)
		return
	}

	productMu.Lock()
	defer productMu.Unlock()
	for i := range products {
		if products[i].ID == id {
			products[i] = input
			utils.WriteJSON(w, http.StatusOK, "success", "Produk berhasil diubah", input)
			return
		}
	}
	utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, ok := productIDFromPath(w, r)
	if !ok {
		return
	}

	if configs.DB != nil {
		result, err := configs.DB.Exec("DELETE FROM products WHERE id = ?", id)
		if err != nil && err != sql.ErrNoRows {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal menghapus produk", nil)
			return
		}
		if affected, _ := result.RowsAffected(); affected == 0 {
			utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
			return
		}
		utils.WriteJSON(w, http.StatusOK, "success", "Produk berhasil dihapus", nil)
		return
	}

	productMu.Lock()
	defer productMu.Unlock()
	for i := range products {
		if products[i].ID == id {
			products = append(products[:i], products[i+1:]...)
			utils.WriteJSON(w, http.StatusOK, "success", "Produk berhasil dihapus", nil)
			return
		}
	}
	utils.WriteJSON(w, http.StatusNotFound, "error", "Produk tidak ditemukan", nil)
}

func productIDFromPath(w http.ResponseWriter, r *http.Request) (int, bool) {
	idText := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(strings.Trim(idText, "/"))
	if err != nil || id <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "ID produk tidak valid", nil)
		return 0, false
	}
	return id, true
}
