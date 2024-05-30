package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/google/uuid"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`SELECT s.id, s.name, s.age, s.gender, c.name FROM student s LEFT JOIN course c ON c.id = s.course_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID(id string) (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`SELECT s.id, s.name, s.age, s.gender, c.name FROM student s LEFT JOIN course c ON c.id = s.course_id WHERE s.id = $1`, id).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (u *StudentRepo) Create(user model.User) error {
	query := `INSERT INTO student (id, name, age, gender, course_id) VALUES ($1, $2, $3, $4, $5)`
	_, err := u.Db.Exec(query, uuid.NewString(), user.Name, user.Age, user.Gender, user.Course)
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepo) Update(user model.User) error {
	query := `UPDATE student SET name = $1, age = $2, gender = $3, course_id = $4 WHERE id = $5`
	_, err := u.Db.Exec(query, user.Name, user.Age, user.Gender, user.Course, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *StudentRepo) Delete(id string) error {
	query := `DELETE FROM student WHERE id = $1`
	_, err := u.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
