package handlers

// ProductView adalah bentuk data produk yang sudah disiapkan untuk ditampilkan ke template.
type ProductView struct {
	Id        string
	Name      string
	Price     float64
	Stock     int
	IsActive  bool
	CreatedAt string
}

// HomePageData berisi data yang dibutuhkan oleh halaman utama.
type HomePageData struct {
	Products []ProductView
	Error    string
}

// EditPageData berisi data yang dibutuhkan oleh halaman edit produk.
type EditPageData struct {
	Product ProductView
	Error   string
}
