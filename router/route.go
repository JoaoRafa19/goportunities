package router

import (
	"github.com/JoaoRafa19/goplaning/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Opening routes
		v1.GET("/opening", handlers.ShowOpening)
		v1.POST("/opening", handlers.CreateOpening)
		v1.DELETE("/opening", handlers.DeleteOpening)
		v1.PUT("/opening", handlers.UpdateOpening)
		v1.GET("/openings", handlers.ListOpening)

	}

}


func CreateOpening(router *gin.Engine){}