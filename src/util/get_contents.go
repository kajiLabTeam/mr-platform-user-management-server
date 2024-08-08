package util

import (
	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
	"github.com/kajiLabTeam/mr-platform-user-management-server/model"
)

func GetContents(userId string) (common.ContentIds, error) {
	contentIds, err := model.GetContents(userId)
	if err != nil {
		return common.ContentIds{}, err
	}
	return contentIds, nil
}
