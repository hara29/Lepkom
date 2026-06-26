package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"product-inventory/handlers"
	"product-inventory/middlewares"
	"product-inventory/utils"
)

func TestGetProductsWithoutToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	rec := httptest.NewRecorder()

	handler := middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetProducts))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected status %d, got %d", http.StatusUnauthorized, rec.Code)
	}
}

func TestGetProductsWithAPIToken(t *testing.T) {
	token, err := utils.GenerateToken(1, "admin_api", "admin", time.Hour)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/api/products", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	handler := middlewares.AuthMiddleware(http.HandlerFunc(handlers.GetProducts))
	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d: %s", http.StatusOK, rec.Code, rec.Body.String())
	}
}
