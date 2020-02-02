package main

import (
	"fmt"
	"net"
	"github.com/trix80dev/subzero2/util"
)

func main() {
	fmt.Print("chehsen")
	listener, err := net.Listen("tcp", "localhost:6112")



	for {

		conn, err := listener.Accept()

		util

	}



}
