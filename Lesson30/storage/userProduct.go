package postgres

import (
	"database/sql"
	"module/model"
)

type UserProductRepo struct {
	Db *sql.DB
}

func NewUserProductRepo(db *sql.DB) *UserProductRepo{
	return &UserProductRepo{Db : db}
}

func (up *UserProductRepo) CreateUserProduct(userProduct model.UserProduct) error {
	tx , err := up.Db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = up.Db.Exec(`INSERT INTO user_products (user_id, product_id) VALUES($1, $2) `, userProduct.UserId, userProduct.ProductId)

	return err
}

func (up *UserProductRepo) GetUserProductByID(id int) (model.UserProduct, error) {
	var userProduct model.UserProduct
	err := up.Db.QueryRow(`SELECT id, user_id, product_id FROM user_products WHERE id = $1`, id).Scan(&userProduct.Id, &userProduct.UserId, &userProduct.ProductId)
	return userProduct, err
}

func (up *UserProductRepo) GetUserProducts() ([]model.UserProduct, error) {
	rows, err := up.Db.Query(`SELECT id, user_id, product_id FROM user_products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userProducts []model.UserProduct
	for rows.Next() {
		var userProduct model.UserProduct
		err := rows.Scan(&userProduct.Id, &userProduct.UserId, &userProduct.ProductId)
		if err != nil {
			return nil, err
		}
		userProducts = append(userProducts, userProduct)
	}
	return userProducts, nil
}

func (up *UserProductRepo) UpdateUserProduct(userProduct model.UserProduct) error {
	tx, err := up.Db.Begin()
	if err != nil {
		return err
	}

	defer func ()  {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = up.Db.Exec(`UPDATE user_products SET user_id = $1, product_id = $2 WHERE id = $3`, userProduct.UserId, userProduct.ProductId, userProduct.Id)

	return err
}

func (up *UserProductRepo) DeleteUserProduct(id int) error {
	tx, err := up.Db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	_, err = tx.Exec(`DELETE FROM user_products WHERE id = $1`, id)
	return err
}