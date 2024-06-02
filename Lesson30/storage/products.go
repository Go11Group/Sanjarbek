package postgres

import (
	"database/sql"
	"module/model"
)

type ProductsRepo struct {
	Db *sql.DB
}

func NewProductRepo(db *sql.DB) *ProductsRepo {
	return &ProductsRepo{Db: db}
}

func (p *ProductsRepo) CreateProduct(product model.Products) error {
	tr, err := p.Db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT into products (id, name, description, price, stock_quantity) VALUES ($1, $2, $3, $4, $5)`
	_, err = p.Db.Exec(query, product.Id, product.Name, product.Description, product.Price, product.Stock_quantity)
	if err != nil {
		return err
	}

	defer func () {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	} ()

	return nil
}


func (p *ProductsRepo) GetProducts() ([]model.Products, error) {
	tr, err := p.Db.Begin()
	if err != nil {
		return nil, err
	}
	var products []model.Products

	rows, err := p.Db.Query(`
		SELECT p.id, name, description, price, stock_quantity, username 
		FROM products as p
		INNER JOIN user_products as up ON p.id = up.product_id
		INNER JOIN users as u ON up.user_id = u.id;
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product model.Products

		err = rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Stock_quantity)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}
	defer func () {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	} ()

	return products, nil
}

func (p *ProductsRepo) UpdateProduct(product model.Products) error {
	_, err := p.Db.Exec(`
		UPDATE products SET name = $1, description = $2, price = $3, stock_quantity = $4 WHERE id = $5
	`, product.Name, product.Description, product.Price, product.Stock_quantity, product.Id)

	return err
}

func (p *ProductsRepo) DeleteProduct(id int) error {
	_, err := p.Db.Exec(`
		DELETE FROM products WHERE id = $1
	`, id)

	return err
}