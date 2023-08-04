package handlers

import (
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)

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
