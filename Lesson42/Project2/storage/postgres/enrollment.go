package postgres

import (
	"database/sql"
	"fmt"
	"module/Project2/model"
	"module/Project2/replace"
	"strings"
	"time"

	"github.com/google/uuid"
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
	id := uuid.NewString()
	enrollment.EnrollmentId = id
	query := `
		INSERT INTO enrollments(
			enrollment_id,
			user_id,
			course_id)
		VALUES($1, $2, $3)`

	_, err = tr.Exec(query, enrollment.EnrollmentId, enrollment.UserId, enrollment.CourseId)
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
		SELECT enrollment_id, user_id, course_id
		FROM enrollments 
		WHERE enrollment_id = $1 AND deleted_at = 0`, id).Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId)

	if err != nil {
		return nil, err
	}

	return &enrollment, nil
}

func (u *EnrollmentRepo) GetAllEnrollments() ([]model.Enrollments, error) {
	query := `SELECT enrollment_id, user_id, course_id
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
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId)
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

func (u *EnrollmentRepo) GetAllEnrollmentsFiltered(f model.EnrollmentGetAll) ([]model.Enrollments, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT enrollment_id, user_id, course_id
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
		err := rows.Scan(&enrollment.EnrollmentId, &enrollment.UserId, &enrollment.CourseId)
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
func (u *EnrollmentRepo) EmrolUpdate(enrolUpdateFilter model.EnrollmentsUpdate) error {
	var params []string
	var args []interface{}
	query := `
	SELECT id
	FROM ENROLLMENTS
	WHERE delete_at IS NULL AND id = $1
	`

	if err := u.db.QueryRow(query, enrolUpdateFilter.EnrolmentId).Err(); err != nil {
		return fmt.Errorf("enrollments by this id not found: %v", err)
	}

	query = `
	UPDATE enrollments SET 
	`

	if enrolUpdateFilter.UserId != nil {
		params = append(params, fmt.Sprintf(" USER_ID = $%d ", len(args)+1))
		args = append(args, *enrolUpdateFilter.UserId)
	}

	if enrolUpdateFilter.CourseId != nil {
		params = append(params, fmt.Sprintf(" course_id = $%d ", len(args)+1))
		args = append(args, *enrolUpdateFilter.CourseId)
	}

	if enrolUpdateFilter.EnrollmentDate != nil {
		params = append(params, fmt.Sprintf("enrollment_date = $%d", len(args)+1))
		args = append(args, *enrolUpdateFilter.EnrollmentDate)
	}

	params = append(params, fmt.Sprintf("update_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, enrolUpdateFilter.EnrolmentId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND delete_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err := u.db.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	fmt.Println(query)
	return nil
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

func (u *EnrollmentRepo) GetEnrolledUsersByCourseID(courseID string) ([]model.Users, error) {
	query := `
		SELECT u.user_id, u.name, u.email, u.birthday
		FROM users u
		JOIN enrollments e ON u.user_id = e.user_id
		WHERE e.course_id = $1 AND e.deleted_at = 0`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.Users

	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday)
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
