package route

import (
	"awesomeProject/api/handler/ihandler"
	"github.com/gin-gonic/gin"
)

func RouteAnimals(g *gin.RouterGroup, handler ihandler.ForestHandler) {
	g.GET("/", handler.GetAnimals)
	g.POST("/populate", handler.PopulateWithAnimals)

}

func RouteInfo(g *gin.RouterGroup, handler ihandler.ForestHandler) {
	g.GET("/", handler.GetSquare)

}
