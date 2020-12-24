package controllers

import (
	"database/sql"
	"errors"
	"products/core/models"
)

// DBProductInstance type.
type DBProductInstance struct {
	DB *sql.DB
}

// IndexProducts function returns all products that
// corresponds with the passed user_id.
func (instance *DBProductInstance) IndexProducts(userID string) ([]models.Product, error) {
	stmt, err := instance.DB.Prepare("SELECT * FROM products WHERE user_id=$1;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)

	if err != nil {
		return nil, err
	}

	products := []models.Product{}

	for rows.Next() {
		product := models.Product{}

		rows.Scan(
			&product.ID,
			&product.UserID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)

		products = append(products, product)
	}

	return products, nil
}

// AddProduct function add a new product in the database with the
// provided data.
func (instance *DBProductInstance) AddProduct(data map[string]interface{}) (*models.Product, error) {
	stmt, err := instance.DB.Prepare("INSERT INTO products(user_id, name, description, price, quantity) VALUES($1, $2, $3, $4, $5) RETURNING *;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(
		data["user_id"],
		data["name"],
		data["description"],
		data["price"],
		data["quantity"],
	)

	if err != nil {
		return nil, err
	}

	product := &models.Product{}

	for rows.Next() {
		rows.Scan(
			&product.ID,
			&product.UserID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)
	}

	return product, nil
}

// UpdateProduct function updates a product data that corresponds
// with the provided id.
func (instance *DBProductInstance) UpdateProduct(data map[string]interface{}) (*models.Product, error) {
	stmt, err := instance.DB.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5 RETURNING *;")

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(
		data["name"],
		data["description"],
		data["price"],
		data["quantity"],
		data["id"],
	)

	if err != nil {
		return nil, err
	}

	product := &models.Product{}

	for rows.Next() {
		rows.Scan(
			&product.ID,
			&product.UserID,
			&product.Name,
			&product.Description,
			&product.Price,
			&product.Quantity,
		)
	}

	if product.ID == "" {
		return nil, errors.New("not-found")
	}

	return product, nil
}

// DeleteProduct function deletes a function in database that
// corresponds with the provided id.
func (instance *DBProductInstance) DeleteProduct(id string) (bool, error) {
	stmt, err := instance.DB.Prepare("DELETE FROM products WHERE id=$1;")

	if err != nil {
		return false, err
	}

	_, err = stmt.Query(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
