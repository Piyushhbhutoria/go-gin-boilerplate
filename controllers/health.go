package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct{}

// Status godoc
// @Summary Health check endpoint
// @Description Returns the health status of the API
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {string} string "Working!"
// @Router /health [get]
func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}
