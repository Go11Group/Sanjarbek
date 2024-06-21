package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
	"module/replace"
	"strings"

	"github.com/google/uuid"
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
	id := uuid.NewString()

	query := `
		INSERT INTO courses(
			course_id,
			title,
			description)
		VALUES($1, $2, $3)`

	_, err = tr.Exec(query, id, course.Title, course.Description)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create course: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}
	course.CourseId = id

	return &course, nil
}

func (u *CourseRepo) GetCourseByID(id string) (*model.Courses, error) {
	var course model.Courses

	err := u.db.QueryRow(`
		SELECT course_id, title, description
		FROM courses 
		WHERE course_id = $1 AND deleted_at = 0`, id).Scan(&course.CourseId, &course.Title, &course.Description)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (u *CourseRepo) GetAllCourses() ([]model.Courses, error) {
	query := `SELECT course_id, title, description
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
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description)
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

func (u *CourseRepo) GetAllCoursesFiltered(f model.CourseGetAll) ([]model.Courses, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT course_id, title, description
	          FROM courses 
	          WHERE deleted_at = 0`

	if f.Title != "" {
		params["title"] = f.Title
		filter += " AND title = :title "
	}

	if f.Description != "" {
		params["description"] = f.Description
		filter += " AND description = :description "
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
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description)
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
	// First, check if the course exists
	var checkCourse model.Courses
	err := u.db.QueryRow(`
		SELECT course_id, title, description
		FROM courses
		WHERE course_id = $1 AND deleted_at = 0`, course.CourseId).Scan(&checkCourse.CourseId, &checkCourse.Title, &checkCourse.Description)
	fmt.Println(err, checkCourse)
	if err != nil {
		return nil, err
	}

	var fields []string
	var args []interface{}
	argID := 1

	if course.Title != "" {
		fields = append(fields, fmt.Sprintf("title = $%d", argID))
		args = append(args, course.Title)
		argID++
	}
	if course.Description != "" {
		fields = append(fields, fmt.Sprintf("description = $%d", argID))
		args = append(args, course.Description)
		argID++
	}

	// If nothing is to be updated, return an error
	if len(fields) == 0 {
		return nil, errors.New("nothing to update")
	}

	query := fmt.Sprintf(`
		UPDATE courses SET
			%s,
			updated_at = CURRENT_TIMESTAMP
		WHERE course_id = $%d AND deleted_at = 0`, 
		strings.Join(fields, ", "),
		argID)

	args = append(args, course.CourseId)

	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	_, err = tr.Exec(query, args...)
	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	// Return the updated course
	updatedCourse, err := u.GetCourseByID(course.CourseId)
	if err != nil {
		return nil, err
	}

	return updatedCourse, nil
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

func (u *CourseRepo) GetMostPopularCourses(startDate, endDate string) ([]model.CoursePopularity, error) {
	query := `
		SELECT c.course_id, c.title, c.description, c.created_at, COUNT(e.user_id) AS enrollment_count
		FROM courses c
		JOIN enrollments e ON c.course_id = e.course_id
		WHERE e.enrollment_date >= $1 AND e.enrollment_date <= $2 AND e.deleted_at = 0
		GROUP BY c.course_id, c.title, c.description, c.created_at
		ORDER BY enrollment_count DESC`

	rows, err := u.db.Query(query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []model.CoursePopularity

	for rows.Next() {
		course := model.CoursePopularity{}
		err := rows.Scan(&course.CourseId, &course.Title, &course.Description, &course.CreatedAt, &course.EnrollmentCount)
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


func (h *CourseRepo) GetCoursesbyUser(id string) ([]model.GetCoursesbyUsers, error) {
	rows, err := h.db.Query("select u.id, c.id,c.title,c.Description from users as u join enrollments as e on u.id=e.user_id join courses as c on c.id=e.course_id where u.id=$1", id)
	if err != nil {
	  panic(err)
	}
  
	var p []model.GetCoursesbyUsers
	var get model.GetCoursesbyUsers
	for rows.Next() {
	  err := rows.Scan(&get.Id, &get.Course.Id, &get.Course.Title, &get.Course.Description)
	  if err != nil {
		return nil, err
	  }
	  p = append(p, get)
	}
	return p, nil
  }
  