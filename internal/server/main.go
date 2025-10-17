package main

import (
	"hermes/internal/models"
	"hermes/internal/server/router"
	"log"
)

func main() {
	models.NewDB()
	e := router.Router()
	err := e.Run()
	if err != nil {
		log.Fatalln("run err.", err)
		return
	} // listen and serve on 0.0.0.0:8080
}
