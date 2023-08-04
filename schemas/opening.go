package schemas

import (
	"time"

	"gorm.io/gorm"
)

const (
	REMOTE = "REMOTE"
	HIBRID = "HIBRID"
	PRESENCIAL = "PRESENCIAL"
)

type Opening struct {
	gorm.Model
	Role      string
	Company   string
	Location  string
	WorkModel string
	Link      string
	Salary    int64
}

type OpeningResponse struct {
	ID         uint      `json:"id"`
	Role       string    `json:"role"`
	Company    string    `json:"company"`
	Location   string    `json:"location"`
	WorkModel  string    `json:"workmodel"`
	Link       string    `json:"link"`
	Salary     int64     `json:"salary"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}
