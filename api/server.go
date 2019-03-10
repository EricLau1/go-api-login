package api

import (
	"net/http"
	"log"
	"fmt"
	"go-api-login/api/routes"
	"go-api-login/api/models"
)

func Run() {
	db := models.Connect()
	if !db.HasTable(&models.User{}) {
		db.Debug().CreateTable(&models.User{})
	}
	db.Close()
	listen(9000)
}

func listen(p int) {
	port := fmt.Sprintf(":%d", p)
	fmt.Printf("Listening Port %s...\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(port, routes.LoadCors(r)))
}