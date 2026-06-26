package handlers

import (
	"net/http"
	"sync"
	"time"

	"product-inventory/configs"
	"product-inventory/middlewares"
	"product-inventory/models"
	"product-inventory/utils"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if configs.DB != nil {
		rows, err := configs.DB.Query(`
			SELECT st.id, st.user_id, COALESCE(u.username, ''), st.product_id, COALESCE(p.name, ''),
			       st.transaction_type, st.quantity, st.created_at
			FROM stock_transactions st
			LEFT JOIN users u ON u.id = st.user_id
			LEFT JOIN products p ON p.id = st.product_id
			ORDER BY st.id DESC`)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal mengambil transaksi", nil)
			return
		}
		defer rows.Close()

		var result []models.StockTransaction
		for rows.Next() {
			var trx models.StockTransaction
			if err := rows.Scan(&trx.ID, &trx.UserID, &trx.Username, &trx.ProductID, &trx.ProductName, &trx.TransactionType, &trx.Quantity, &trx.CreatedAt); err != nil {
				utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membaca transaksi", nil)
				return
			}
			result = append(result, trx)
		}
		utils.WriteJSON(w, http.StatusOK, "success", "Daftar transaksi", result)
		return
	}

	transactionMu.Lock()
	defer transactionMu.Unlock()
	utils.WriteJSON(w, http.StatusOK, "success", "Daftar transaksi", transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var input models.StockTransaction
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Body request tidak valid", nil)
		return
	}
	if input.TransactionType != "in" && input.TransactionType != "out" {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "transaction_type harus in atau out", nil)
		return
	}
	if input.ProductID <= 0 || input.Quantity <= 0 {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "product_id dan quantity wajib valid", nil)
		return
	}

	if user, ok := middlewares.UserFromContext(r.Context()); ok {
		input.UserID = user.ID
		input.Username = user.Username
	}
	input.CreatedAt = time.Now()

	if configs.DB != nil {
		tx, err := configs.DB.Begin()
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal memulai transaksi", nil)
			return
		}
		defer tx.Rollback()

		delta := input.Quantity
		if input.TransactionType == "out" {
			delta = -delta
		}
		if _, err := tx.Exec("UPDATE products SET stock = stock + ? WHERE id = ? AND stock + ? >= 0", delta, input.ProductID, delta); err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal memperbarui stok", nil)
			return
		}
		result, err := tx.Exec(
			"INSERT INTO stock_transactions (user_id, product_id, transaction_type, quantity) VALUES (?, ?, ?, ?)",
			input.UserID,
			input.ProductID,
			input.TransactionType,
			input.Quantity,
		)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal menyimpan transaksi", nil)
			return
		}
		if err := tx.Commit(); err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal menyelesaikan transaksi", nil)
			return
		}
		id, _ := result.LastInsertId()
		input.ID = int(id)
		utils.WriteJSON(w, http.StatusCreated, "success", "Transaksi berhasil dibuat", input)
		return
	}

	transactionMu.Lock()
	input.ID = nextTransactionID
	nextTransactionID++
	transactions = append(transactions, input)
	transactionMu.Unlock()
	utils.WriteJSON(w, http.StatusCreated, "success", "Transaksi berhasil dibuat", input)
}

var (
	transactionMu     sync.Mutex
	transactions      []models.StockTransaction
	nextTransactionID = 1
)
