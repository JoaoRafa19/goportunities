package handlers

import (
	"fmt"
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a Job opening description
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]
func ShowOpening(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == ""	{
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	opening := schemas.Opening{}

	// Find
	if err := db.First(&opening, id).Error ; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return 
	}
	
	sendSuccess(ctx, "show-opening", opening)
}