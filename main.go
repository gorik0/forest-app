package main

import (
	"awesomeProject/config"
	"awesomeProject/di"
)

func main() {

	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	server, err := di.InitServer()
	if err != nil {
		panic(err)

	}

	err = server.Run()
	if err != nil {
		panic(err)
	}

	println("HElo gorik!!!")

}
