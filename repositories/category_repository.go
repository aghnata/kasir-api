package repositories

import (
	"database/sql"
	"errors"
	"kasir-api/models"
	// "errors"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

func (repo *CategoryRepository) GetAll() ([]models.Category, error) {
	query := "SELECT id, name, description FROM categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *CategoryRepository) Create(category *models.Category) error {
	// Validate input
	if category == nil {
		return sql.ErrNoRows
	}
	// Use parameterized query to prevent SQL injection
	query := "INSERT INTO categories (name, description) VALUES ($1, $2) RETURNING id"
	err := repo.db.QueryRow(query, category.Name, category.Description).Scan(&category.ID)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CategoryRepository) GetByID(id int) (*models.Category, error) {
	query := "SELECT id, name, description FROM categories WHERE id = $1"
	row := repo.db.QueryRow(query, id)
	var category models.Category
	err := row.Scan(&category.ID, &category.Name, &category.Description)

	if err == sql.ErrNoRows {
		return nil, errors.New("category tidak ditemukan")
	}
	if err != nil {
		return nil, err
	}
	return &category, nil
}
