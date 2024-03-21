package controller

import (
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{"status": statusCode, "message": errorMessage})
}

func SendSuccessResponse(c *gin.Context, statusCode int, successMessage string) {
	c.JSON(statusCode, gin.H{"status": statusCode, "message": successMessage})
}
