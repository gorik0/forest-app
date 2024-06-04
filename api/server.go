package api

import (
	"awesomeProject/api/handler/ihandler"
	"awesomeProject/api/route"
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	engine *gin.Engine
	addr   AddrString
}

type AddrString string

func ProvideAddrString(config *config.Config) AddrString {
	return AddrString(config.HttpAddr)
}
func NewHttpServer(forestHandler ihandler.ForestHandler, address AddrString) *HttpServer {
	s := HttpServer{engine: gin.New(), addr: address}
	route.RouteAnimals(s.engine.Group("/animals"), forestHandler)
	route.RouteInfo(s.engine.Group("/info"), forestHandler)

	return &s
}

func (s *HttpServer) Run() error {
	return s.engine.Run(string(s.addr))
}
