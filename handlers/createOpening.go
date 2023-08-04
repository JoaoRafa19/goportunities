package handlers

import (
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Sumary Create opening
// @Description Create a new Job opening
// @Tags Opening
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		logger.ErrF("validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	logger.InfoF("request received: %+v", request)

	opening := schemas.Opening{
		Role:      request.Role,
		Company:   request.Company,
		Location:  request.Location,
		Link:      request.Link,
		Salary:    request.Salary,
		WorkModel: request.WorkModel,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrF("Erro ao criar vaga: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error on create opening")
		return
	}
	sendSuccess(ctx, "createopening", opening)

}
