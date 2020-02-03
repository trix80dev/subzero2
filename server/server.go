package server

import (
	"net"
	"yukon/packet"
	"yukon/util"
)

const (
	LOGIN      = 0
	GAME       = 1
	REDEMPTION = 2
)

type YukonServer struct {
	Name       string
	Port       string
	ServerType byte
}

func NewServer(name string, serverType byte, port string) *YukonServer {
	s := new(YukonServer)
	s.Port = port
	s.ServerType = serverType
	s.Name = name
	return s
}

func (s YukonServer) StartServer() {

	listener, err := net.Listen("tcp", "localhost:" + s.Port)
	util.HandleError(err)

	for {

		conn, err := listener.Accept()
		util.HandleError(err)

		//fmt.Println("[" + "] Client connected to server " + s.Name)

		if s.ServerType == LOGIN {
			go packet.HandleLogin(conn)
		}

	}

}
