package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"module/model"
	"module/replace"
)

type LessonRepo struct {
	db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{db: db}
}

func (u *LessonRepo) CreateLesson(lesson model.Lessons) (*model.Lessons, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `
		INSERT INTO lessons(
			lesson_id,
			course_id,
			title,
			content)
		VALUES($1, $2, $3, $4)`

	_, err = tr.Exec(query, lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create lesson: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (u *LessonRepo) GetLessonByID(id string) (*model.Lessons, error) {
	var lesson model.Lessons

	err := u.db.QueryRow(`
		SELECT lesson_id, course_id, title, content 
		FROM lessons 
		WHERE lesson_id = $1 AND deleted_at = 0`, id).Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content)

	if err != nil {
		return nil, err
	}

	return &lesson, nil
}

func (u *LessonRepo) GetAllLessons() (*[]model.Lessons, error) {
	query := `SELECT lesson_id, course_id, title, content 
	          FROM lessons 
	          WHERE deleted_at = 0`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons

	for rows.Next() {
		lesson := model.Lessons{}
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &lessons, nil
}

func (u *LessonRepo) GetAllLessonsFiltered(f model.LessonGetAll) ([]model.Lessons, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT lesson_id, course_id, title, content 
	          FROM lessons 
	          WHERE deleted_at = 0 `

	if f.Title != "" {
		params["title"] = f.Title
		filter += " AND title = :title "
	}

	if f.Content != "" {
		params["content"] = f.Content
		filter += " AND content = :content "
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

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lessons, nil
}

func (u *LessonRepo) UpdateLesson(lesson model.Lessons) (*model.Lessons, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	res, err := tr.Exec(`
		UPDATE lessons SET
			course_id = $2,
			title = $3,
			content = $4
		WHERE lesson_id = $1 AND deleted_at = 0`, lesson.LessonId, lesson.CourseId, lesson.Title, lesson.Content)

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

	return &lesson, nil
}

func (u *LessonRepo) DeleteLesson(id string) error {
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(`
		UPDATE lessons SET
 		deleted_at = date_part('epoch', current_timestamp)::BIGINT
		WHERE lesson_id = $1 AND deleted_at = 0`, id)

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
