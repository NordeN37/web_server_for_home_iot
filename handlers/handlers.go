package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"remote_commands_for_home_IoT/depends"
	"remote_commands_for_home_IoT/handlers/room_controller"
	"remote_commands_for_home_IoT/middleware"
)

func MakeServiceHandler(di depends.IAppDeps, r *mux.Router) {
	r.Use(middleware.New(di.Config().TokenUsers).Middleware)
	r.HandleFunc("/turn_lamp_in_room", room_controller.NewRoomController(di).TurnLampInRoom).Methods(http.MethodPost)
}
