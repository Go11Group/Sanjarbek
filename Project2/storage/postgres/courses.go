package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
	"module/replace"
)

type CourseRepo struct {
	db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{db: db}
}

func (u *CourseRepo) CreateCourse(course model.Courses) (*model.Courses, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO courses(
			course_id,
			title,
			description,
			created_at)
		VALUES($1, $2, $3, $4)`

	_, err = tr.Exec(query, course.CourseId, course.Title, course.Description, course.CreatedAt)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create course: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (u *CourseRepo) GetCourseByID(id string) (*model.Courses, error) {
	var course model.Courses

	err := u.db.QueryRow(`
		SELECT course_id, title, description, created_at 
		FROM courses 
		WHERE course_id = $1 AND deleted_at = 0`, id).Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (u *CourseRepo) GetAllCourses() (*[]model.Courses, error) {
	query := `SELECT course_id, title, description, created_at 
	          FROM courses 
	          WHERE deleted_at = 0`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.Courses

	for rows.Next() {
		course := model.Courses{}
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &courses, nil
}

func (u *CourseRepo) GetAllCoursesFiltered(f model.CourseGetAll) ([]model.Courses, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT course_id, title, description, created_at 
	          FROM courses 
	          WHERE deleted_at = 0 `

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " AND title = :title "
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

	var courses []model.Courses
	for rows.Next() {
		var course model.Courses
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt)
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

func (u *CourseRepo) UpdateCourse(course model.Courses) (*model.Courses, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	res, err := tr.Exec(`
		UPDATE courses SET
			title = $2,
			description = $3,
			updated_at = CURRENT_TIMESTAMP
		WHERE course_id = $1 AND deleted_at = 0`, course.CourseId, course.Title, course.Description)
	
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

	return &course, nil
}

func (u *CourseRepo) DeleteCourse(id string) error {
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(`
		UPDATE courses SET
 		deleted_at = date_part('epoch', current_timestamp)::BIGINT
		WHERE course_id = $1 AND deleted_at = 0`, id)
	
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
