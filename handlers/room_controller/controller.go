package room_controller

import "remote_commands_for_home_IoT/depends"

type RoomController struct {
	di depends.IAppDeps
}

func NewRoomController(di depends.IAppDeps) *RoomController {
	return &RoomController{
		di: di,
	}
}
