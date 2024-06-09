package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(user model.User) (*model.User, error) {
	query := `
		INSERT INTO users(
			id,
			first_name,
			last_name,
			field,
			email)
		VALUES($1, $2, $3, $4, $5)`

	_, err := u.db.Exec(query, user.Id, user.FirstName, user.LastName, user.Field, user.Email)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	return &user, nil
}

func (u *UserRepo) GetUserByID(id int) (model.User, error) {
	var user model.User

	err := u.db.QueryRow(`
		SELECT id, first_name, last_name, field, email FROM users WHERE id = $1
	`, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Field, &user.Email)

	return user, err
}

func (u *UserRepo) GetAll() (*[]model.User, error) {
	query := `SELECT id, first_name, last_name, field, email FROM users`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Field,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &users, err
}

func (u *UserRepo) Update(user model.User) (*model.User, error) {
	res, err := u.db.Exec(`
		UPDATE users SET
			first_name = $2,
			last_name = $3,
			field = $4,
			email = $5
		WHERE id = $1`, user.Id, user.FirstName, user.LastName, user.Field, user.Email)
	
	if err != nil {
		return nil, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if n == 0 {
		return nil, errors.New("this id is not available")
	}

	return &user, err
}

func (u *UserRepo) Delete(id int) error {
	_, err := u.db.Exec(`
	DELETE FROM users WHERE id = $1`, id)
	
	return err
}
