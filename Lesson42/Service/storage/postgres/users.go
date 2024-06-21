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

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) Create(user model.Users) (*model.Users, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}
	id := uuid.NewString()

	query := `
		INSERT INTO users(
			user_id,
			name,
			email,
			birthday,
			password)
		VALUES($1, $2, $3, $4, $5)`

	_, err = tr.Exec(query, id, user.Name, user.Email, user.Birthday, user.Password)
	if err != nil {
		tr.Rollback()
		return nil, fmt.Errorf("could not create user: %v", err)
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}
	user.UserId = id

	return &user, nil
}

func (u *UserRepo) GetUserByID(id string) (*model.Users, error) {
	var user model.Users
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	err = u.db.QueryRow(`
		SELECT user_id, name, email, birthday, password
		FROM users 
		WHERE user_id = $1 AND deleted_at = 0`, id).Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password)

	if err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) GetAllUsers() (*[]model.Users, error) {
	tr, err := u.db.Begin()
	if err != nil {
		return nil, err
	}

	query := `SELECT user_id, name, email, birthday, password
	          FROM users 
	          WHERE deleted_at = 0`

	rows, err := u.db.Query(query)
	if err != nil {
		tr.Rollback()
		return nil, err
	}
	defer rows.Close()

	var users []model.Users

	for rows.Next() {
		user := model.Users{}
		err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password)
		if err != nil {
			tr.Rollback()
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		tr.Rollback()
		return nil, err
	}

	err = tr.Commit()
	if err != nil {
		return nil, err
	}

	return &users, nil
}

func (u *UserRepo) GetAllUsersFiltered(f model.UserGetAll) ([]model.Users, error) {
	var (
	  params = make(map[string]interface{})
	  arr    []interface{}
	  filter string
	)
  
	query := `SELECT user_id, name, email, birthday, password
	FROM users WHERE true `
	fmt.Println(f)
  
	if len(f.Name) > 0 {
	  params["name"] = f.Name
	  filter += " and name = :name "
	}
  
	if f.Birthday != "" {
	  params["birthday"] = f.Birthday
	  filter += " and birthday = :birthday "
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
	fmt.Println(query, arr)
	rows, err := u.db.Query(query, arr...)
	fmt.Println(err)
	if err != nil {
	  return nil, err
	}
  
	var users []model.Users
	for rows.Next() {
	  var user model.Users
	  err := rows.Scan(&user.UserId, &user.Name, &user.Email, &user.Birthday, &user.Password)
  
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

func (u *UserRepo) Update(user model.Users, id string) (*model.Users, error) {
	// Birinchi user mavjudligini tekshiramiz
	var checkUser model.Users
	err := u.db.QueryRow(`
		SELECT user_id, name, email, birthday, password
		FROM users 
		WHERE user_id = $1 AND deleted_at = 0`, id).Scan(&checkUser.UserId, &checkUser.Name, &checkUser.Email, &checkUser.Birthday, &checkUser.Password)

	if err != nil {
		return nil, err
	}

	var fields []string
	var args []interface{}
	argID := 1

	if user.Name != "" {
		fields = append(fields, fmt.Sprintf("name = $%d", argID))
		args = append(args, user.Name)
		argID++
	}
	if user.Email != "" {
		fields = append(fields, fmt.Sprintf("email = $%d", argID))
		args = append(args, user.Email)
		argID++
	}
	if user.Birthday != "" {
		fields = append(fields, fmt.Sprintf("birthday = $%d", argID))
		args = append(args, user.Birthday)
		argID++
	}
	if user.Password != "" {
		fields = append(fields, fmt.Sprintf("password = $%d", argID))
		args = append(args, user.Password)
		argID++
	}

	// Agar hech narsa o'zgartirilmasa, xatolikni qaytaramiz
	if len(fields) == 0 {
		return nil, errors.New("nothing to update")
	}

	query := fmt.Sprintf(`
		UPDATE users SET
			%s,
			updated_at = CURRENT_TIMESTAMP
		WHERE user_id = $%d AND deleted_at = 0`,
		strings.Join(fields, ", "),
		argID)

	args = append(args, id)

	_, err = u.db.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	// Yangi holatni qaytaramiz
	updatedUser, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (u *UserRepo) Delete(id string) error {
	tr, err := u.db.Begin()
	if err != nil {
		return err
	}

	_, err = tr.Exec(`
		UPDATE users SET
 		deleted_at = date_part('epoch', current_timestamp)::BIGINT
		WHERE user_id = $1 AND deleted_at = 0`, id)

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

func (u *UserRepo) SearchUsers(fSearch model.UsersGet) (*model.Result, error) {
	var (
	  params = make(map[string]interface{})
	  args   []interface{}
	  filter string
	)
  
	query := "SELECT user_id, name, email FROM users WHERE deleted_at = 0 "
  
	if fSearch.Name != "" {
	  params["name"] = fSearch.Name
	  filter += "AND name = :name "
	}
	if fSearch.Email != "" {
	  params["email"] = fSearch.Email
	  filter += "AND email = :email "
	}
	if fSearch.AgeTo > 0 && fSearch.AgeFrom > 0 {
	  params["age_to"] = fSearch.AgeTo
	  params["age_from"] = fSearch.AgeFrom
  
	  filter += "AND EXTRACT(YEAR FROM age(birthday)) between :age_from AND :age_to "
	}
  
	query += filter
  
	query, args = replace.ReplaceQueryParams(query, params)
  
	rows, err := u.db.Query(query, args...)
	if err != nil {
	  return nil, err
	}
  
	var results []model.AdditionalUser
	for rows.Next() {
	  var result model.AdditionalUser
  
	  err = rows.Scan(&result.ID, &result.Name, &result.Email)
  
	  if err != nil {
		return nil, err
	  }
  
	  results = append(results, result)
	}
  
	if err := rows.Err(); err != nil {
	  return nil, err
	}
  
	return &model.Result{Results: results}, err
  }