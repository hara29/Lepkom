package models

import "time"

// Product merepresentasikan struktur data produk beserta detailnya dari database.
type Product struct {
	Id        string
	Name      string
	Price     float64
	Stock     int
	IsActive  bool
	CreatedAt time.Time
	Image     []byte
}
