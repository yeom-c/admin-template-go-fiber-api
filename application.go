package main

import (
	"github.com/yeom-c/admin-template-go-fiber-api/app"
	"github.com/yeom-c/admin-template-go-fiber-api/middleware"
	"github.com/yeom-c/admin-template-go-fiber-api/route"
	"log"
)

func main() {
	middleware.SetMiddleware()
	route.SetRoutes()

	err := app.Server().Run()
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
