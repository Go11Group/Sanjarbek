package postgres

import (
	"database/sql"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/google/uuid"
)

type CourseRepo struct {
	Db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{Db: db}
}

func (r *CourseRepo) GetAllCourses() ([]model.Course, error) {
	rows, err := r.Db.Query(`SELECT id, name, field FROM course`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Course
	for rows.Next() {
		var course model.Course
		err = rows.Scan(&course.Id, &course.Name, &course.Field)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (r *CourseRepo) GetCourseByID(id string) (model.Course, error) {
	var course model.Course

	err := r.Db.QueryRow(`SELECT id, name, field FROM course WHERE id = $1`, id).
		Scan(&course.Id, &course.Name, &course.Field)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func (r *CourseRepo) CreateCourse(course model.Course) error {
	query := `INSERT INTO course (id, name, field) VALUES ($1, $2, $3)`
	_, err := r.Db.Exec(query, uuid.New(), course.Name, course.Field)
	if err != nil {
		return err
	}
	return nil
}

func (r *CourseRepo) UpdateCourse(course model.Course) error {
	query := `UPDATE course SET name = $1, field = $2 WHERE id = $3`
	_, err := r.Db.Exec(query, course.Name, course.Field, course.Id)
	if err != nil {
		return err
	}
	return nil
}

func (r *CourseRepo) DeleteCourse(id string) error {
	query := `DELETE FROM course WHERE id = $1`
	_, err := r.Db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
