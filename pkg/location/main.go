package location

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
)

const botName = "ChatCord Bot"

type roomDto struct {
	Username string `json:"username"`
	Room     string `json:"room"`
}

type roomInfo struct {
	Room  string `json:"room"`
	Users []User `json:"users"`
}

func InitChatEndpoints(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		return nil
	})

	server.OnEvent("/", "join-room", func(s socketio.Conn, dto roomDto) {
		user := User{Id: s.ID(), Username: dto.Username, Room: dto.Room}

		s.SetContext(dto)

		s.Join(user.Room)

		users = append(users, user)

		server.BroadcastToRoom(
			"/",
			user.Room,
			"message",
			FormatMessage(botName, fmt.Sprintf(`%s has joined the chat`, user.Username)),
		)

		server.BroadcastToRoom(
			"/",
			user.Room,
			"room-users",
			roomInfo{
				Room:  user.Room,
				Users: GetRoomUsers(user.Room),
			},
		)
	})

	server.OnEvent("/", "chat-message", func(s socketio.Conn, msg string) {
		user, err := GetUser(s.ID())

		if err != nil {
			log.Println(err.Error())
			return
		}

		server.BroadcastToRoom(
			"/",
			user.Room,
			"message",
			FormatMessage(user.Username, msg),
		)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		user, err := GetUser(s.ID())

		if err != nil {
			log.Println(err.Error())
			return
		}

		RemoveUser(s.ID())

		server.BroadcastToRoom(
			"/",
			user.Room,
			"message",
			FormatMessage(botName, fmt.Sprintf(`%s has left the chat`, user.Username)),
		)

		server.BroadcastToRoom(
			"/",
			user.Room,
			"room-users",
			roomInfo{
				Room:  user.Room,
				Users: GetRoomUsers(user.Room),
			},
		)

		s.Leave(user.Room)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
}
