package model

import (
	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
)

func InsertContent(userId string, createdId string, contentId string) (bool, error) {
	uuid := uuid.New()
	_, err := db.Exec("INSERT INTO user_contents (id,user_id,created_id,content_id) VALUES ($1,$2,$3,$4)", uuid.String(), userId, createdId, contentId)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GetContents(userId string) (common.ContentIds, error) {
	// 最新のcreated_idを持つ content_id を取得
	rows, err := db.Query("SELECT content_id FROM user_contents WHERE user_id = $1 ORDER BY created_id DESC LIMIT 1", userId)
	if err != nil {
		return common.ContentIds{}, err
	}
	defer rows.Close()

	var contentIds common.ContentIds
	for rows.Next() {
		var contentId string
		if err := rows.Scan(&contentId); err != nil {
			return common.ContentIds{}, err
		}
		contentIds.ContentIds = append(contentIds.ContentIds, contentId)
	}
	return contentIds, nil
}
