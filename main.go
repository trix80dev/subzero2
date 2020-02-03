package main

import (
	"fmt"
	"yukon/server"
)

func main() {
	fmt.Println("Yukon started")

	loginServer := server.NewServer("Login", server.LOGIN, "6112")
	loginServer.StartServer()

}
