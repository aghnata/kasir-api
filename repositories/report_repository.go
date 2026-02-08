package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (repo *ReportRepository) GenerateSalesReport() (interface{}, error) {
	// Implement report generation logic here
	query := `select 
	p."name" product_name,
	count(t.id ) total_transaction,
	sum(td.quantity ) as total_qty,
	sum(td.subtotal) as total_revenue
	from transactions t 
	left join transaction_details td on td.transaction_id = t.id 
	left join products p on td.product_id = p.id 
	where t.created_at::date = now()::date
	group by 
		p."name" 
	order by total_qty desc`

	// args := []any{}
	// if nameFilter != "" {
	// 	query += " WHERE p.name ILIKE $1"
	// 	args = append(args, "%"+nameFilter+"%")
	// }

	// rows, err := repo.db.Query(query, args...)
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]models.Report, 0)
	totalRevenue := 0
	totalTransactions := 0
	totalQty := 0
	for rows.Next() {
		var report models.Report
		err := rows.Scan(&report.ProductName, &report.TotalTransaction, &report.TotalQty, &report.TotalRevenue)
		if err != nil {
			return nil, err
		}
		totalQty += report.TotalQty
		totalRevenue += report.TotalRevenue
		totalTransactions += report.TotalTransaction
		result = append(result, report)
	}

	fmt.Println(result)

	resultDto := models.SalesReport{
		TotalRevenue:      totalRevenue,
		TotalTransactions: totalTransactions,
		BestSellerProduct: models.BestSellerProduct{
			Name:         result[0].ProductName,
			TotalQtySold: result[0].TotalQty,
		},
	}

	return resultDto, nil
}
