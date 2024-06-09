package postgres

import (
	"database/sql"
	"errors"
	"module/model"
)

type SolvedProblemRepo struct {
	db *sql.DB
}

func NewSolvedProblemRepo(db *sql.DB) *SolvedProblemRepo {
	return &SolvedProblemRepo{db: db}
}

func (p *SolvedProblemRepo) Create(solvedProblem model.SolvedProblems) (*model.SolvedProblems, error) {
	_, err := p.db.Exec(`solvedProblem
		INSERT into problem(
			id, user_id, name, difficulty, explanation) VALUES($1, $2, $3, $4, $5)
			`, solvedProblem.Id, solvedProblem.UserId, solvedProblem.Name, solvedProblem.Difficulty, solvedProblem.Explanation)

		
	if err != nil {
		return nil, err
	}

	return &solvedProblem, err
}

func (p *SolvedProblemRepo) GetById(id int) (*model.SolvedProblems, error) {
	solvedProblem := model.SolvedProblems{}
	err := p.db.QueryRow(`
		SELECT * from problem WHERE id = $1
		`, id).Scan(&solvedProblem.Id, &solvedProblem.UserId, &solvedProblem.Name, &solvedProblem.Difficulty, &solvedProblem.Explanation)

	if err != nil {
		return nil, err
	}

	return &solvedProblem, err
}

func (p *SolvedProblemRepo) GetAll() (*[]model.SolvedProblems, error) {
	query := `SELECT * FROM problem`

	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}

	var solvedProblems []model.SolvedProblems

	for rows.Next() {
		solvedProblem := model.SolvedProblems{}
		rows.Scan(
			&solvedProblem.Id,
			&solvedProblem.UserId,
			&solvedProblem.Name,
			&solvedProblem.Difficulty,
			&solvedProblem.Explanation,
		)
		solvedProblems = append(solvedProblems, solvedProblem)
	}

	err = rows.Close()
	if err != nil {
		return nil, err
	}

	return &solvedProblems, err
}

func (p *SolvedProblemRepo) Update(solvedProblem model.SolvedProblems) (*model.SolvedProblems, error) {
	res, err := p.db.Exec(`
		UPDATE problem SET
		SET
			user_id = $2.
			name = $3,
			difficulty = $4,
			explanation = $5
		WHERE id = $1`, solvedProblem.Id, solvedProblem.UserId, solvedProblem.Name, solvedProblem.Difficulty, solvedProblem.Explanation)
	
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

	return &solvedProblem, err
}

func (p *SolvedProblemRepo) Delete(id int) error {
	_, err := p.db.Exec(`
	DELETE FROM problems WHERE id = $1`, id)
	
	return err
}