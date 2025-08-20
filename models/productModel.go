package models

type Produk struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Harga     float64 `json:"harga"`
	Kategori  string  `json:"kategori"`
	Stok      int     `json:"stok"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
