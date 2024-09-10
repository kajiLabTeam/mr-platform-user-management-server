package model

import (
	"database/sql"
)

func ExistUser(userId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM users WHERE id = $1", userId)

	var user string
	if err := row.Scan(&user); err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned, return false and no error
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func CreateUser(userId string) (bool, error) {
	_, err := db.Exec("INSERT INTO users (id) VALUES ($1)", userId)
	if err != nil {
		return false, err
	}
	return true, nil
}
