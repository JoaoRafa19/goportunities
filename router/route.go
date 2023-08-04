package router

import (
	"github.com/JoaoRafa19/goplaning/docs"
	"github.com/JoaoRafa19/goplaning/handlers"
	"github.com/gin-gonic/gin"
	swagfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler 
	handlers.Init()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		// Opening routes
		v1.GET("/opening", handlers.ShowOpening)
		v1.POST("/opening", handlers.CreateOpening)
		v1.DELETE("/opening", handlers.DeleteOpening)
		v1.PUT("/opening", handlers.UpdateOpening)
		v1.GET("/openings", handlers.ListOpening)

	}
	// Init Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagfiles.Handler))

}


func CreateOpening(router *gin.Engine){}