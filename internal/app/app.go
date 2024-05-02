package app

import (
	"InterviewTest/internal/config"
	"InterviewTest/internal/handler"
	"InterviewTest/internal/server"
	"InterviewTest/internal/service"
)

func Run(configPath string) {
	//config
	configContainer, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	//handler
	servicesContainer := service.NewService()
	handlerProcessor := handler.NewHandler(servicesContainer)

	//server
	serverInstance := server.NewServer(configContainer.Port, handlerProcessor)
	err = serverInstance.Run()
	if err != nil {
		panic(err)
	}
}
