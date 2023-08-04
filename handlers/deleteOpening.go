package handlers

import (
	"fmt"
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Delete opening
// @Description Delete a new Job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [delete]
func DeleteOpening(ctx *gin.Context) {
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
	
	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s not found", id))
		return 
	}
	sendSuccess(ctx, "delete-opening", opening)
}