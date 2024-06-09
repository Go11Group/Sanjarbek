package postgres

import (
	"database/sql"
	"errors"
	"module/model"

	"github.com/google/uuid"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(user model.User) (*model.User, error) {
	user.Id = uuid.NewString()

	_, err := u.db.Exec(`\
		INSERT INTO user(
		id,
		first_name,
		lasr_name,
		field,
		email)
		VALUES($1, $2, $3, $4, $5)`, user.Id, user.FirstName, user.LastName, user.Field, user.Email)

	if err != nil {
		return nil, err
	}

	return &user, err
}

func (u *UserRepo) GetById(id string) (model.User, error) {
	user := model.User{}
	err := u.db.QueryRow(`
	SELECT * FROM user WHERE id = $1`, id).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Field, &user.Email)

	if err != nil {
		return user, err
	}

	return user, err
}

func (u *UserRepo) GetAll() (*[]model.User, error) {
	query := `SELECT * FROM problem`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}

	var users []model.User

	for rows.Next() {
		user := model.User{}
		rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.Field,
			&user.Email,
		)
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
		UPDATE problem SET
		SET
			name = $2,
			difficulty = $3,
			explanation = $4
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

func (u *UserRepo) Delete(id string) error {
	_, err := u.db.Exec(`
	DELETE FROM problems WHERE id = $1`, id)
	
	return err
}