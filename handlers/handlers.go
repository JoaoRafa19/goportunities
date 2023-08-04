package handlers

import (
	"github.com/JoaoRafa19/goplaning/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db *gorm.DB  
)

func Init(){
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
