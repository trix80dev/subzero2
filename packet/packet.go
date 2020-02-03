package packet

import "net"

type Packet struct {
	data string
	conn net.Conn
}
