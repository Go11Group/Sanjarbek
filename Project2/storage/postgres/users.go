package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
	"module/replace"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(user model.Users) (*model.Users, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO users(
			name,
			email,
			birthday,
			password)
		VALUES($1, $2, $3, $4)`

	err = tr.QueryRow(query, user.Name, user.Email, user.Birthday, user.Password).Scan(&user.UserId, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetUserByID(id string) (*model.Users, error) {
	var user model.Users

	err := u.db.QueryRow(`
		SELECT user_id, name, email, birthday, password, created_at, updated_at 
		FROM users 
		WHERE user_id = $1 AND deleted_at = 0`, id).Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAll() (*[]model.Users, error) {
	query := `SELECT user_id, name, email, birthday, password, created_at, updated_at 
	          FROM users 
	          WHERE deleted_at = 0`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users

	for rows.Next() {
		user := model.Users{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepo) GetAllUsers(f model.UserGetAll) ([]model.Users, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT user_id, name, email, birthday, password 
	          FROM users 
	          WHERE deleted_at = 0 `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " AND name = :name "
	}

	if f.Birthday != "" {
		params["birthday"] = f.Birthday
		filter += " AND birthday = :birthday "
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		filter += " OFFSET :offset"
	}

	if f.Limit > 0 {
		params["limit"] = f.Limit
		filter += " LIMIT :limit"
	}

	query = query + filter

	query, arr = replace.ReplaceQueryParams(query, params)
	rows, err := u.db.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) Update(user model.Users) (*model.Users, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	res, err := tr.Exec(`
		UPDATE users SET
			name = $2,
			email = $3,
			birthday = $4,
			password = $5,
			updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $1 AND deleted_at = 0`, user.UserId, user.Name, user.Email, user.Birthday, user.Password)

	if err != nil {
		tr.Rollback()
		return nil, err
	}

	n, err := res.RowsAffected()
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	if n == 0 {
		tr.Rollback()
		return nil, errors.New("this id is not available")
	}

	return &user, nil
}

func (u *UserRepo) Delete(id string) error {
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(`
		UPDATE users SET
 		deleted_at = date_part('epoch', current_timestamp)::BIGINT
		WHERE user_id = $1 AND deleted_at = 0`, id)

	if err != nil {
		tr.Rollback()
		return err
	}

	err = tr.Commit()
	if err != nil {
		return err
	}

	return nil
}
