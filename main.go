package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"quizmo/apis"
	"quizmo/core"
	"quizmo/utils"

	"github.com/labstack/echo/v4/middleware"
)

func main() {

	serverConfig := &utils.Config{}
	configBytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(configBytes, serverConfig)

	if err != nil {
		log.Fatal(err)
	}
	db := utils.NewDatabase(serverConfig)
	server := core.NewServer(serverConfig, db)
	// Add middlewares
	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())
	server.Echo.Use(middleware.CORS())

	_ = apis.Movies(server)
	server.Start()
}
