package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
	"github.com/kajiLabTeam/mr-platform-user-management-server/model"
)

func SetContents(c *gin.Context) {
	var req common.RequestSetContents
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdId, err := uuid.NewV7()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	isExistUser, err := model.IsExistUser(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !isExistUser {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	for _, contentId := range req.ContentIds {
		isInserted, err := model.InsertContent(req.UserId, createdId.String(), contentId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if !isInserted {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert content"})
			return
		}
	}

	contenIds, err := model.GetContents(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, contenIds)
}

// CREATE USER mr_platform_contents WITH PASSWORD 'mr_platform_contents';
// CREATE USER mr_platform_users WITH PASSWORD 'mr_platform_users';

// \c mr_platform;

// -- Create Users Table
// CREATE TABLE users (
//     id VARCHAR(36) PRIMARY KEY,
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE,
//     deleted_at TIMESTAMP WITH TIME ZONE
// );

// -- Create User Contents Table
// CREATE TABLE user_contents (
//     id VARCHAR(36) PRIMARY KEY,
//     user_id VARCHAR(36) REFERENCES users(id),
//     created_id VARCHAR(36) NOT NULL,
//     content_id VARCHAR(36) REFERENCES contents(id),
//     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
//     updated_at TIMESTAMP WITH TIME ZONE,
//     deleted_at TIMESTAMP WITH TIME ZONE
// );

// GRANT ALL PRIVILEGES ON TABLE users TO mr_platform_users;
// GRANT ALL PRIVILEGES ON TABLE user_contents TO mr_platform_users;

