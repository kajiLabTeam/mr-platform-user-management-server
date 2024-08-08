package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kajiLabTeam/mr-platform-user-management-server/common"
	"github.com/kajiLabTeam/mr-platform-user-management-server/util"
)

func GetContents(c *gin.Context) {
	var req common.RequestGetContents
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contentIds, err := util.GetContents(req.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, contentIds)
}
