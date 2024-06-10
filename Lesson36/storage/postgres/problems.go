package postgres

import (
	"database/sql"
	"errors"
	"module/model"
)

type ProblemRepo struct {
	db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{db: db}
}

func (p *ProblemRepo) Create(problem model.Problems) (*model.Problems, error) {
	_, err := p.db.Exec(`
		INSERT into problems(
			id, name, difficulty, explanation) VALUES($1, $2, $3, $4)
			`, problem.Id, problem.Name, problem.Difficulty, problem.Explanation)

		
	if err != nil {
		return nil, err
	}

	return &problem, err
}

func (p *ProblemRepo) GetById(id int) (*model.Problems, error) {
	problem := model.Problems{}
	err := p.db.QueryRow(`
		SELECT * from problem WHERE id = $1
		`, id).Scan(&problem.Id, &problem.Name, &problem.Difficulty, &problem.Explanation)

	if err != nil {
		return nil, err
	}

	return &problem, err
}

func (p *ProblemRepo) GetAll() (*[]model.Problems, error) {
	query := `SELECT * FROM problems`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	var problems []model.Problems

	for rows.Next() {
		problem := model.Problems{}
		rows.Scan(
			&problem.Id,
			&problem.Name,
			&problem.Difficulty,
			&problem.Explanation,
		)
		problems = append(problems, problem)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &problems, err
}

func (p *ProblemRepo) Update(problem model.Problems) (*model.Problems, error) {
	res, err := p.db.Exec(`
		UPDATE problems SET
		SET
			name = $2,
			difficulty = $3,
			explanation = $4
		WHERE id = $1`, problem.Id, problem.Name, problem.Difficulty, problem.Explanation)
	
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

	return &problem, err
}

func (p *ProblemRepo) Delete(id int) error {
	_, err := p.db.Exec(`
	DELETE FROM problems WHERE id = $1`, id)
	
	return err
}