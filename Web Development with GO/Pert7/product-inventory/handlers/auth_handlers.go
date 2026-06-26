package handlers

import (
	"database/sql"
	"net/http"
	"strings"
	"sync"
	"time"

	"product-inventory/configs"
	"product-inventory/models"
	"product-inventory/utils"
)

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var (
	userMu     sync.Mutex
	users      = []models.User{{ID: 1, Username: "admin", Password: "password123", Role: "admin"}, {ID: 2, Username: "user1", Password: "password123", Role: "user"}}
	nextUserID = 3
)

func Login(w http.ResponseWriter, r *http.Request) {
	var input credentials
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Body request tidak valid", nil)
		return
	}

	user, err := findUserByCredentials(input.Username, input.Password)
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized, "error", "Username atau password salah!", nil)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Role, 2*time.Hour)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membuat token", nil)
		return
	}
	utils.WriteJSON(w, http.StatusOK, "success", "Login berhasil", map[string]interface{}{"token": token, "user": user})
}

func Register(w http.ResponseWriter, r *http.Request) {
	var input credentials
	if err := utils.DecodeJSON(r, &input); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Body request tidak valid", nil)
		return
	}
	input.Username = strings.TrimSpace(input.Username)
	if input.Username == "" || input.Password == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Username dan password wajib diisi", nil)
		return
	}
	if input.Role == "" {
		input.Role = "user"
	}
	if input.Role != "admin" && input.Role != "user" {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Role harus admin atau user", nil)
		return
	}

	user := models.User{Username: input.Username, Password: input.Password, Role: input.Role}
	if configs.DB != nil {
		result, err := configs.DB.Exec("INSERT INTO users (username, password, role) VALUES (?, ?, ?)", user.Username, user.Password, user.Role)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membuat user", nil)
			return
		}
		id, _ := result.LastInsertId()
		user.ID = int(id)
		user.Password = ""
		utils.WriteJSON(w, http.StatusCreated, "success", "User berhasil dibuat", user)
		return
	}

	userMu.Lock()
	user.ID = nextUserID
	nextUserID++
	users = append(users, user)
	user.Password = ""
	userMu.Unlock()

	utils.WriteJSON(w, http.StatusCreated, "success", "User berhasil dibuat", user)
}

func GenerateAPIToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
	}
	_ = utils.DecodeJSON(r, &input)
	input.Username = strings.TrimSpace(input.Username)
	if input.Username == "" {
		input.Username = strings.TrimSpace(r.URL.Query().Get("username"))
	}
	if input.Username == "" {
		utils.WriteJSON(w, http.StatusBadRequest, "error", "Username wajib diisi", nil)
		return
	}

	user, err := findUserByUsername(input.Username)
	if err != nil {
		utils.WriteJSON(w, http.StatusNotFound, "error", "User tidak ditemukan", nil)
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Username+"_api", user.Role, 24*365*time.Hour)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membuat token", nil)
		return
	}
	if configs.DB != nil {
		_, _ = configs.DB.Exec("UPDATE users SET api_token = ? WHERE id = ?", token, user.ID)
	}

	user.APIToken = token
	saveFallbackAPIToken(user.ID, token)
	utils.WriteJSON(w, http.StatusOK, "success", "Token API berhasil dibuat", map[string]string{"username": user.Username, "role": user.Role, "api_token": token})
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	if configs.DB != nil {
		rows, err := configs.DB.Query("SELECT id, username, role, COALESCE(api_token, '') FROM users ORDER BY id")
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal mengambil user", nil)
			return
		}
		defer rows.Close()

		var result []models.User
		for rows.Next() {
			var user models.User
			if err := rows.Scan(&user.ID, &user.Username, &user.Role, &user.APIToken); err != nil {
				utils.WriteJSON(w, http.StatusInternalServerError, "error", "Gagal membaca user", nil)
				return
			}
			result = append(result, user)
		}
		utils.WriteJSON(w, http.StatusOK, "success", "Daftar user", result)
		return
	}

	userMu.Lock()
	defer userMu.Unlock()
	var result []models.User
	for _, user := range users {
		user.Password = ""
		result = append(result, user)
	}
	utils.WriteJSON(w, http.StatusOK, "success", "Daftar user", result)
}

func findUserByCredentials(username, password string) (models.User, error) {
	if configs.DB != nil {
		var user models.User
		err := configs.DB.QueryRow(
			"SELECT id, username, password, role, COALESCE(api_token, '') FROM users WHERE username = ? AND password = ?",
			username,
			password,
		).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.APIToken)
		user.Password = ""
		return user, err
	}

	userMu.Lock()
	defer userMu.Unlock()
	for _, user := range users {
		if user.Username == username && user.Password == password {
			user.Password = ""
			return user, nil
		}
	}
	return models.User{}, sql.ErrNoRows
}

func findUserByUsername(username string) (models.User, error) {
	if configs.DB != nil {
		var user models.User
		err := configs.DB.QueryRow(
			"SELECT id, username, password, role, COALESCE(api_token, '') FROM users WHERE username = ?",
			username,
		).Scan(&user.ID, &user.Username, &user.Password, &user.Role, &user.APIToken)
		user.Password = ""
		return user, err
	}

	userMu.Lock()
	defer userMu.Unlock()
	for _, user := range users {
		if user.Username == username {
			user.Password = ""
			return user, nil
		}
	}
	return models.User{}, sql.ErrNoRows
}

func saveFallbackAPIToken(userID int, token string) {
	userMu.Lock()
	defer userMu.Unlock()
	for i := range users {
		if users[i].ID == userID {
			users[i].APIToken = token
			return
		}
	}
}
