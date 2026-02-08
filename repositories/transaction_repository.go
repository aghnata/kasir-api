package repositories

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	// Collect all product IDs
	productIDs := make([]any, len(items))
	itemMap := make(map[int]models.CheckoutItem)
	for i, item := range items {
		productIDs[i] = item.ProductID
		itemMap[item.ProductID] = item
	}

	// Build dynamic query for IN clause
	query := "SELECT id, name, price, stock FROM products WHERE id IN ("
	for i := range productIDs {
		if i > 0 {
			query += ", "
		}
		query += fmt.Sprintf("$%d", i+1)
	}
	query += ")"

	// Fetch all products in a single query
	rows, err := tx.Query(query, productIDs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productData := make(map[int]struct {
		name  string
		price int
		stock int
	})

	for rows.Next() {
		var id, price, stock int
		var name string
		if err := rows.Scan(&id, &name, &price, &stock); err != nil {
			return nil, err
		}
		productData[id] = struct {
			name  string
			price int
			stock int
		}{name, price, stock}
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Validate all products exist
	for _, item := range items {
		if _, exists := productData[item.ProductID]; !exists {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
	}

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	// Calculate totals and prepare details
	for _, item := range items {
		product := productData[item.ProductID]
		subtotal := product.price * item.Quantity
		totalAmount += subtotal

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: product.name,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	// Batch update stock using a single query with CASE
	updateQuery := "UPDATE products SET stock = stock - CASE id "
	updateArgs := make([]any, 0, len(items)*2)
	argIdx := 1
	for _, item := range items {
		updateQuery += fmt.Sprintf("WHEN $%d THEN $%d ", argIdx, argIdx+1)
		updateArgs = append(updateArgs, item.ProductID, item.Quantity)
		argIdx += 2
	}
	updateQuery += "ELSE 0 END WHERE id IN ("
	for i, item := range items {
		if i > 0 {
			updateQuery += ", "
		}
		updateQuery += fmt.Sprintf("$%d", argIdx)
		updateArgs = append(updateArgs, item.ProductID)
		argIdx++
		_ = item
	}
	updateQuery += ")"

	_, err = tx.Exec(updateQuery, updateArgs...)
	if err != nil {
		return nil, err
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	// Batch insert transaction details using a single query
	insertQuery := "INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES "
	insertArgs := make([]any, 0, len(details)*4)
	argIdx = 1
	for i, detail := range details {
		if i > 0 {
			insertQuery += ", "
		}
		insertQuery += fmt.Sprintf("($%d, $%d, $%d, $%d)", argIdx, argIdx+1, argIdx+2, argIdx+3)
		insertArgs = append(insertArgs, transactionID, detail.ProductID, detail.Quantity, detail.Subtotal)
		details[i].TransactionID = transactionID
		argIdx += 4
	}

	_, err = tx.Exec(insertQuery, insertArgs...)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}
