package handlers

import (
	"fmt"
	"net/http"

	"github.com/JoaoRafa19/goplaning/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrF("Validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParam").Error())
	}

	opening := schemas.Opening{}

	// Find
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Update opening

	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.WorkModel != "" {
		opening.WorkModel = request.WorkModel
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.ErrF("error update opening", err)
		sendError(ctx, http.StatusInternalServerError, "error updating opening")
		return
	}
	sendSuccess(ctx, "update-opening", opening)

}

/*
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
*/
