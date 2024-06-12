package postgres

import (
	"database/sql"
	"fmt"
	"module/model"
)

type UserRepository interface {
	Create(user model.Users) (model.Users, error)
	GetUserByID(id string) (model.Users, error)
	GetAll() ([]model.Users, error)
	Update(user model.Users) (model.Users, error)
	Delete(id string) error
}

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user model.Users) (model.Users, error) {
	_, err := r.db.Exec("INSERT INTO users (name, email, birthday, password) VALUES ($1, $2, $3, $4)",
		user.Name, user.Email, user.Birthday, user.Password)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to create user: %v", err)
	}
	return user, nil
}

func (r *UserRepo) GetUserByID(id string) (model.Users, error) {
	var user model.Users
	row := r.db.QueryRow("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users WHERE user_id = $1", id)
	err := row.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Users{}, fmt.Errorf("user not found")
		}
		return model.Users{}, fmt.Errorf("failed to get user: %v", err)
	}
	return user, nil
}

func (r *UserRepo) GetAll() ([]model.Users, error) {
	var users []model.Users
	rows, err := r.db.Query("SELECT user_id, name, email, birthday, password, created_at, updated_at, deleted_at FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user row: %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over users rows: %v", err)
	}
	return users, nil
}

func (r *UserRepo) Update(user model.Users) (model.Users, error) {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, birthday = $3, password = $4, updated_at = CURRENT_TIMESTAMP WHERE user_id = $5",
		user.Name, user.Email, user.Birthday, user.Password, user.UserId)
	if err != nil {
		return model.Users{}, fmt.Errorf("failed to update user: %v", err)
	}
	return user, nil
}

func (r *UserRepo) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE user_id = $1", id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
