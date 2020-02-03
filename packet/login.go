package packet

import (
	"bytes"
	"fmt"
	"net"
	"yukon/util"
)

var XMLPackets = map[string]func(conn net.Conn, packet string) string{
	"<policy-file-request/>": policyFileRequest,
}

func HandleLogin(conn net.Conn) {
	for {

		data := make([]byte, 64)

		read, err := conn.Read(data)
		if read > 0 {
			util.HandleError(err)

			data = bytes.Trim(data, "\x00")

			inPacket := string(data)

			fmt.Println("IN: " + inPacket)

			outPacket := XMLPackets[inPacket](conn, inPacket) + "\u0000"
			conn.Write([]byte(outPacket))
			fmt.Println("OUT: " + outPacket)
		}
	}
}

func policyFileRequest(conn net.Conn, packet string) string {
	return "<cross-domain-policy><allow-access-from domain='*' to-ports='" + conn.LocalAddr().String()[len(conn.LocalAddr().String()) - 4:] + "'/></cross-domain-policy>"
}
