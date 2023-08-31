package room_controller

import (
	"net/http"
	"remote_commands_for_home_IoT/mqtt"
)

func (rc RoomController) TurnLampInRoom(w http.ResponseWriter, _ *http.Request) {
	rc.di.MQTT().MQTTPublish(mqtt.TopicRoom, []byte("asdasdasd"))
	w.WriteHeader(200)
}
