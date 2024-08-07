package model

import (
	"database/sql"

	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
)

func IsExistUser(userId string) (bool, error) {
	row := db.QueryRow("SELECT id FROM users WHERE id = $1", userId)

	var user common.UserId
	if err := row.Scan(&user.UserId); err != nil {
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
