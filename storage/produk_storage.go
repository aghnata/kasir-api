package storage

import (
	"errors"
	"kasir-api/models"
)

// ProdukStorage handles product data operations
type ProdukStorage struct {
	produk []models.Produk
}

// NewProdukStorage creates a new ProdukStorage instance with initial data
func NewProdukStorage() *ProdukStorage {
	return &ProdukStorage{
		produk: []models.Produk{
			{ID: 1, Nama: "Produk A", Harga: 10000, Stok: 50},
			{ID: 2, Nama: "Produk B", Harga: 20000, Stok: 30},
			{ID: 3, Nama: "Produk C", Harga: 15000, Stok: 20},
		},
	}
}

// GetAll returns all products
func (ps *ProdukStorage) GetAll() []models.Produk {
	return ps.produk
}

// GetByID returns a product by ID
func (ps *ProdukStorage) GetByID(id int) (*models.Produk, error) {
	for _, p := range ps.produk {
		if p.ID == id {
			return &p, nil
		}
	}
	return nil, errors.New("produk tidak ditemukan")
}

// Create adds a new product
func (ps *ProdukStorage) Create(produkBaru models.Produk) models.Produk {
	produkBaru.ID = len(ps.produk) + 1
	ps.produk = append(ps.produk, produkBaru)
	return produkBaru
}

// Update updates an existing product by ID
func (ps *ProdukStorage) Update(id int, produkUpdate models.Produk) (*models.Produk, error) {
	for i, p := range ps.produk {
		if p.ID == id {
			produkUpdate.ID = id
			ps.produk[i] = produkUpdate
			return &produkUpdate, nil
		}
	}
	return nil, errors.New("produk tidak ditemukan")
}

// Delete removes a product by ID
func (ps *ProdukStorage) Delete(id int) error {
	for i, p := range ps.produk {
		if p.ID == id {
			ps.produk = append(ps.produk[:i], ps.produk[i+1:]...)
			return nil
		}
	}
	return errors.New("produk tidak ditemukan")
}
