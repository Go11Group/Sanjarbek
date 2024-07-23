package postgres

import (
	"database/sql"
	"gin/api/models"
)

type NewAuth struct {
	Db *sql.DB
}

func NewAuthRepo(db *sql.DB) *NewAuth {
	return &NewAuth{Db: db}
}

func (u *NewAuth) Register(req *models.AuthReq) error {

	_, err := u.Db.Exec(`INSERT INTO
							users(name, user_name, password, email)
						VALUES($1, $2, $3, $4)`,
		req.Name, req.UserName, req.Password, req.Email)

	if err != nil {
		return err
	}
	return nil

}

func (u *NewAuth) Login(req *models.LoginReq) (*models.UserRes, error) {
	user := models.UserRes{}
	err := u.Db.QueryRow(`SELECT id, name, user_name, email
       FROM users
       WHERE user_name = $1 and passwotd = $2`, req.UserName, req.Password).Scan(user.Id, user.Name, user.UserName, user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil

}
