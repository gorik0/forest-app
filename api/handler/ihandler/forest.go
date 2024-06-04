package ihandler

import "github.com/gin-gonic/gin"

type ForestHandler interface {
	GetAnimals(c *gin.Context)
	GetSquare(c *gin.Context)
	PopulateWithAnimals(c *gin.Context)
}
