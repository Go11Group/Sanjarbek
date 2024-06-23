package postgres

import (
	"database/sql"
	"user_service/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Create(user *models.User) error {
	id := uuid.NewString()
	_, err := u.Db.Exec(`
        INSERT INTO users(id, name, phone, age)
        VALUES($1, $2, $3, $4)`,
		id, user.Name, user.Phone, user.Age)
	return err
}

func (u *UserRepo) GetById(id string) (*models.User, error) {
	user := &models.User{Id: id}
	err := u.Db.QueryRow(`
        SELECT name, phone, age
        FROM users
        WHERE id = $1`, id).Scan(&user.Name, &user.Phone, &user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) GetAll() ([]models.User, error) {
	rows, err := u.Db.Query(`
        SELECT id, name, phone, age
        FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		user := models.User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Phone, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepo) Update(user *models.User) (*models.User, error) {
	id := user.Id
	var checkUser models.User
	err := u.Db.QueryRow(`
		SELECT id, name, phone, age
		FROM users
		WHERE id = $1`, id).Scan(&checkUser.Id, &checkUser.Name, &checkUser.Phone, &checkUser.Age)
	if err != nil {
		return nil, err
	}

	if user.Id == "" {
		user.Id = checkUser.Id
	}
	if user.Name == "" {
		user.Name = checkUser.Name
	}
	if user.Phone == "" {
		user.Phone = checkUser.Phone
	}
	if user.Age == 0 {
		user.Age = checkUser.Age
	}

	_, err = u.Db.Exec(`
        UPDATE users
        SET name = $2, phone = $3, age = $4
        WHERE id = $1`,
		user.Id, user.Name, user.Phone, user.Age)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepo) Delete(id string) error {
	_, err := u.Db.Exec(`DELETE FROM users WHERE id = $1`, id)
	return err
}
