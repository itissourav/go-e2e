package db

import (
	"database/sql"
	"go-e2e/models.go"
)

func ListUsers(db *sql.DB) ([]models.User, error) {

	query := `
		SELECT id, first_name,last_name, email
		FROM users
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.Id,
			&user.Firstname,
			&user.Lastname,
			&user.Email,
		)
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

func CreateUser(db *sql.DB, user models.SignupReq) error {

	query := `
		INSERT INTO public.users (
			first_name,
			last_name,
			email,
			password
		)
		VALUES ($1, $2, $3, $4)
	`

	_, err := db.Exec(
		query,
		user.Firstname,
		user.Lastname,
		user.Email,
		user.Password,
	)

	return err
}

func UserExists(db *sql.DB, email string) (bool, error) {

	var exists bool

	query := `
		SELECT EXISTS(
			SELECT 1
			FROM users
			WHERE email = $1
		)
	`

	err := db.QueryRow(query, email).Scan(&exists)

	return exists, err
}
