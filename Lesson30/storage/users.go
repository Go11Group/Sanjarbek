package postgres

import (
	"database/sql"
	"module/model"
)

type UsersRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{Db: db}
}

func (u *UsersRepo) CreateUser(user model.Users) error {
	tr, err := u.Db.Begin()
	if err != nil {
		return err
	}
	_, err = u.Db.Exec(`
		INSERT INTO users(id, username, email, password) 
		VALUES
			($1, $2, $3, $4)
	`, user.Id, user.UserName, user.Email, user.Password)

	defer func () {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	} ()

	return err
}

func (u *UsersRepo) GetUsers() ([]model.Users, error) {
	tr, err := u.Db.Begin()
	if err != nil {
		return nil, err
	}

	var users []model.Users

	rows, err := u.Db.Query(`
		SELECT u.id, username, email, password, name
		FROM users as u 
		INNER JOIN user_products as up ON u.id = up.user_id
		INNER JOIN products as p ON up.product_id = p.id
	`)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user model.Users

		err = rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	defer func () {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	} ()

	return users, nil
}

func (u *UsersRepo) UpdateUser(user model.Users) error {
	_, err := u.Db.Exec(`UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4`, user.UserName, user.Email, user.Password, user.Id)
	return err
}


func (u *UsersRepo) DeleteUser(id int) error {
	_, err := u.Db.Exec(`DELETE FROM users WHERE id = $1`, id)

	return err
}