package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
	"github.com/kajiLabTeam/mr-platform-user-management-server/model"
)

func CreateUser(c *gin.Context) {
	var user common.UserId
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユーザーが存在するか確認
	exist, err := model.ExistUser(user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if exist {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}

	// ユーザーを作成
	isCreated, err := model.CreateUser(user.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isCreated {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
