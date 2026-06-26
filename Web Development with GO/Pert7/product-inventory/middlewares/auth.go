package middlewares

import (
	"context"
	"database/sql"
	"net/http"
	"strings"

	"product-inventory/configs"
	"product-inventory/models"
	"product-inventory/utils"
)

type contextKey string

const userContextKey contextKey = "auth_user"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")
		if authorization == "" {
			utils.WriteJSON(w, http.StatusUnauthorized, "error", "Mana Tokennya?", nil)
			return
		}

		token, err := utils.BearerToken(authorization)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, "error", err.Error(), nil)
			return
		}

		user, err := userFromToken(token, r)
		if err != nil {
			utils.WriteJSON(w, http.StatusUnauthorized, "error", "Token tidak valid", nil)
			return
		}

		ctx := context.WithValue(r.Context(), userContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func AdminOnly(next http.Handler) http.Handler {
	return AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, ok := UserFromContext(r.Context())
		if !ok || user.Role != "admin" {
			utils.WriteJSON(w, http.StatusForbidden, "error", "Akses ditolak: hanya admin yang diizinkan", nil)
			return
		}
		next.ServeHTTP(w, r)
	}))
}

func UserFromContext(ctx context.Context) (models.User, bool) {
	user, ok := ctx.Value(userContextKey).(models.User)
	return user, ok
}

func userFromToken(token string, r *http.Request) (models.User, error) {
	if claims, err := utils.ValidateToken(token); err == nil {
		username := claims.Username
		if strings.HasSuffix(username, "_api") {
			username = strings.TrimSuffix(username, "_api")
			if err := ensureStoredAPIToken(username, token); err != nil {
				return models.User{}, err
			}
		}

		if r.Header.Get("Referer") == "" && r.Header.Get("Origin") == "" && !strings.HasSuffix(claims.Username, "_api") {
			return models.User{}, sql.ErrNoRows
		}

		return models.User{
			ID:       claims.UserID,
			Username: username,
			Role:     claims.Role,
			APIToken: token,
		}, nil
	}

	return userByStoredAPIToken(token)
}

func ensureStoredAPIToken(username, token string) error {
	if configs.DB == nil {
		return nil
	}

	var id int
	return configs.DB.QueryRow("SELECT id FROM users WHERE username = ? AND api_token = ?", username, token).Scan(&id)
}

func userByStoredAPIToken(token string) (models.User, error) {
	if configs.DB == nil {
		return models.User{}, sql.ErrNoRows
	}

	var user models.User
	err := configs.DB.QueryRow(
		"SELECT id, username, role, COALESCE(api_token, '') FROM users WHERE api_token = ?",
		token,
	).Scan(&user.ID, &user.Username, &user.Role, &user.APIToken)
	return user, err
}
