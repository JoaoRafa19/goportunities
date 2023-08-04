package handlers

import (
	"fmt"

	"github.com/JoaoRafa19/goplaning/schemas"
)

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

//Create Opening

type CreateOpeningRequest struct {
	Role      string `json:"role"`
	Company   string `json:"company"`
	Location  string `json:"location"`
	WorkModel string `json:"workmodel"`
	Link      string `json:"link"`
	Salary    int64  `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	if r.WorkModel == "" {
		return errParamIsRequired("workmodel", "string")
	}
	if r.WorkModel != schemas.REMOTE && r.WorkModel != schemas.HIBRID && r.WorkModel != schemas.PRESENCIAL {
		return errParamIsRequired("workmodel", "string")
	}
	if r.Salary <= 0 {
		return errParamIsRequired("salary", "int64")
	}

	return nil
}
