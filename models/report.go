package models

type Report struct {
	ProductName      string `json:"product_name"`
	TotalTransaction int    `json:"total_transaction"`
	TotalQty         int    `json:"total_qty"`
	TotalRevenue     int    `json:"total_revenue"`
}

type SalesReport struct {
	TotalRevenue      int               `json:"total_revenue"`
	TotalTransactions int               `json:"total_transaksi"`
	BestSellerProduct BestSellerProduct `json:"produk_terlaris"`
}

type BestSellerProduct struct {
	Name         string `json:"nama"`
	TotalQtySold int    `json:"qty_terjual"`
}
