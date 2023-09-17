package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omfj/lol/internal/models"
	"github.com/omfj/lol/internal/security"
)

func GetStatus(c *gin.Context) {
	status, err := models.GetStatus()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/plain", []byte("failed to get status"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}

func UpdateStatus(c *gin.Context) {
	status := c.Query("status")

	if status == "" {
		c.Data(http.StatusBadRequest, "text/plain", []byte("status query parameter is required"))
		return
	}

	token := c.GetHeader("Authorization")
	if token == "" {
		c.Data(http.StatusUnauthorized, "text/plain", []byte("authorization header is required"))
		return
	}

	token = strings.Split(token, " ")[1]

	err := security.ValidateToken(token)
	if err != nil {
		c.Data(http.StatusUnauthorized, "text/plain", []byte("invalid token"))
		return
	}

	err = models.SetStatus(status)
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/plain", []byte("failed to set status"))
		return
	}

	status, err = models.GetStatus()
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/plain", []byte("failed to get status"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": status,
	})
}
