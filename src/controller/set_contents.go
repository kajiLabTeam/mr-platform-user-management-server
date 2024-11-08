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

	existUser, err := model.ExistUser(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !existUser {
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

	res := common.ResponseSetContents{
		ContentIds: contenIds,
	}

	c.JSON(http.StatusCreated, res)
}
