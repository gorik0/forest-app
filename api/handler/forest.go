package handler

import (
	"awesomeProject/api/handler/ihandler"
	"awesomeProject/usecase/iservice"
	"awesomeProject/utils/response"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type ForestHandler struct {
	s      iservice.ForestService
	logger *slog.Logger
}

func (f *ForestHandler) GetAnimals(c *gin.Context) {
	const op = "api.handler.GetAnimals"

	logger := f.logger.With(slog.String("op", op))
	f.logger.Debug("Starting")
	animals, err := f.s.GetAnimals()

	if err != nil {
		logger.Error("error while geting animals")
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "coudln't get animals ", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessReponse("Gotta animals from forests", animals))
	return
}

func (f *ForestHandler) GetSquare(c *gin.Context) {
	square, err := f.s.GetSquare()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse(http.StatusInternalServerError, "Fail to get square", err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessReponse("GOTTA square", square))
	return
}

func (f *ForestHandler) PopulateWithAnimals(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func NewForestHandler(f iservice.ForestService, logger *slog.Logger) ihandler.ForestHandler {
	return &ForestHandler{f, logger}
}
