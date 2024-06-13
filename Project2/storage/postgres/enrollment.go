package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
	"module/replace"
)

type EnrollmentRepo struct {
	db *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{db: db}
}

func (u *EnrollmentRepo) CreateEnrollment(enrollment model.Enrollments) (*model.Enrollments, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO enrollments(
			enrollment_id,
			user_id,
			course_id,
			enrollment_date)
		VALUES($1, $2, $3, $4)`

	_, err = tr.Exec(query, enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create enrollment: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (u *EnrollmentRepo) GetEnrollmentByID(id string) (*model.Enrollments, error) {
	var enrollment model.Enrollments

	err := u.db.QueryRow(`
		SELECT enrollment_id, user_id, course_id, enrollment_date 
		FROM enrollments 
		WHERE enrollment_id = $1 AND deleted_at = 0`, id).Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)

	if err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (u *EnrollmentRepo) GetAllEnrollments() (*[]model.Enrollments, error) {
	query := `SELECT enrollment_id, user_id, course_id, enrollment_date 
	          FROM enrollments 
	          WHERE deleted_at = 0`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var enrollments []model.Enrollments

	for rows.Next() {
		enrollment := model.Enrollments{}
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &enrollments, nil
}

func (u *EnrollmentRepo) GetAllEnrollmentsFiltered(f model.EnrollmentGetAll) ([]model.Enrollments, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT enrollment_id, user_id, course_id, enrollment_date 
	          FROM enrollments 
	          WHERE deleted_at = 0 `

	if f.UserId != "" {
		params["user_id"] = f.UserId
		filter += " AND user_id = :user_id "
	}

	if f.CourseId != "" {
		params["course_id"] = f.CourseId
		filter += " AND course_id = :course_id "
	}

	if f.EnrollmentDate != "" {
		params["enrollment_date"] = f.EnrollmentDate
		filter += " AND enrollment_date = :enrollment_date "
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

	var enrollments []model.Enrollments
	for rows.Next() {
		var enrollment model.Enrollments
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId, &enrollment.EnrollmentDate)
		if err != nil {
			return nil, err
		}
		enrollments = append(enrollments, enrollment)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return enrollments, nil
}

func (u *EnrollmentRepo) UpdateEnrollment(enrollment model.Enrollments) (*model.Enrollments, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	res, err := tr.Exec(`
		UPDATE enrollments SET
			user_id = $2,
			course_id = $3,
			enrollment_date = $4
		WHERE enrollment_id = $1 AND deleted_at = 0`, enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId, enrollment.EnrollmentDate)

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

	return &enrollment, nil
}

func (u *EnrollmentRepo) DeleteEnrollment(id string) error {
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(`
		UPDATE enrollments SET
 		deleted_at = date_part('epoch', current_timestamp)::BIGINT
		WHERE enrollment_id = $1 AND deleted_at = 0`, id)

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
