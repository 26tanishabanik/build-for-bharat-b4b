package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// handler for validating whether API is running or not
func HealthCheck(c *gin.Context) {
	now := time.Now()
	health := make(map[string]string)
	health["now"] = now.Format(time.ANSIC)
	c.JSON(http.StatusOK, health)
}
