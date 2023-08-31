package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"remote_commands_for_home_IoT/config"
	"remote_commands_for_home_IoT/depends"
	"remote_commands_for_home_IoT/handlers"
	"remote_commands_for_home_IoT/mqtt"
	"time"
)

func main() {
	cfg := config.New()

	mqttClient, err := mqtt.NewClient(cfg.Endpoint, cfg.ClientID, cfg.User, cfg.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer mqttClient.Close()

	di := depends.NewDI(mqttClient, cfg)

	r := mux.NewRouter()
	handlers.MakeServiceHandler(di, r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
