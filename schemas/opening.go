package schemas

import (
	"gorm.io/gorm"
)

type WorkModel int

const (
	REMOTE     = 1
	PRESENCIAL = 2
	HIBRID     = 3
)

func (s WorkModel) String() string {
	switch s {
	case REMOTE:
		return "REMOTE"
	case PRESENCIAL:
		return "PRESENCIAL"
	case HIBRID:
		return "HIBRID"
	}
	return "unknown"
}

type Opening struct {
	gorm.Model
	Role      string
	Company   string
	Location  string
	WorkModel WorkModel
	Link      string
	Salary    int64
}

