package handlers

import (
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)


// @BasePath /api/v1

// @Summary Show all openigns
// @Description Show all jobs
// @Tags Opening
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpening(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error list openings")
	}
	sendSuccess(ctx, "listopenings", openings)
}
