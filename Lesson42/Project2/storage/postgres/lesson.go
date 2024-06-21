package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"module/Project2/model"
	"module/Project2/replace"
	"strings"

	"github.com/google/uuid"
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
		log.Printf("Error starting transaction: %v", err)
		return nil, err
	}

	// lesson Id yaratadi
	id := uuid.NewString()
	lesson.LessonId = id

	query := `
		INSERT INTO lessons(
			lesson_id,
			course_id,
			title,
			content)
		VALUES($1, $2, $3, $4)`

	_, err = tr.Exec(query, id, lesson.CourseId, lesson.Title, lesson.Content)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create lesson: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		log.Printf("Error committing transaction: %v", err)
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

func (u *LessonRepo) GetAllLessons(f model.LessonGetAll) ([]model.Lessons, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

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
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons
	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.CourseId, &lesson.Title, &lesson.Content)
		if err != nil {
			tr.Rollback()
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err = rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return lessons, nil
}

func (u *LessonRepo) UpdateLesson(lesson model.LessonsUpdate) (*model.LessonsUpdate, error) {
    // Begin a transaction
    tr, err := u.db.Begin()
    if err != nil {
        return nil, err
    }
    defer tr.Commit() // Commit the transaction at the end

    // Prepare the query and arguments
    var args []interface{}
    argID := 1
    var fields []string

    // Check if Title needs to be updated
    if lesson.Title != "" {
        fields = append(fields, fmt.Sprintf("title = $%d", argID))
        args = append(args, lesson.Title)
        argID++
    }

    // Check if Content needs to be updated
    if lesson.Content != "" {
        fields = append(fields, fmt.Sprintf("content = $%d", argID))
        args = append(args, lesson.Content)
        argID++
    }

    // Check if there are any fields to update
    if len(fields) == 0 {
        return nil, errors.New("nothing to update")
    }

    args = append(args, lesson.LessonId) // Assuming LessonId is already set

    query := fmt.Sprintf(`
        UPDATE lessons SET
            %s,
            updated_at = CURRENT_TIMESTAMP
        WHERE lesson_id = $%d AND deleted_at = 0`,
        strings.Join(fields, ", "),
        argID)

    _, err = tr.Exec(query, args...)
    if err != nil {
        tr.Rollback()
        return nil, err
    }

    // Return the updated lesson
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

func (u *LessonRepo) GetLessonsByCourseID(courseID string) ([]model.Lessons, error) {
	query := `
		SELECT lesson_id, title, content
		FROM lessons
		WHERE course_id = $1 AND deleted_at = 0`

	rows, err := u.db.Query(query, courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lessons

	for rows.Next() {
		var lesson model.Lessons
		err := rows.Scan(&lesson.LessonId, &lesson.Title, &lesson.Content)
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