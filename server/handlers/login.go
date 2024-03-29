package handlers

import (
	"log"
	"../jswebsocket"
	"../players"
)

// --------------------------------------------------- //
// Login
// --------------------------------------------------- //

type loginRecv struct {
	Username string `json:username`
	Password string `json:password`
}
type loginResp struct {
}
var loginHandler = func(connection *jswebsocket.Connection, stream interface{}) {
	message := stream.(*loginRecv)
	log.Print("login", message)
	log.Print("user", connection.User)

	// If the connection has a user, error
	if connection.User != nil {
		connection.Message("login", &errorMessage{ Message: "You are already logged in" })
		return
	}

	player := players.Player{
		Name: message.Username,
		Connection: connection,
	}
	player.Register()

	// TODO: not thread safe
	connection.User = &player
	connection.Message("login", &loginResp{})
}