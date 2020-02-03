package packet

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

var XMLPackets = map[string]func(conn net.Conn, packet string) string{
	"<policy-file-request/>": policyFileRequest,
	"<msg t='sys'><body action='verChk' r='0'><ver v='153' /></body></msg>": versionCheck,
	"<msg t='sys'><body action='rndK' r='-1'></body></msg>": randomKey,
}

func HandleLogin(conn net.Conn) {
	defer conn.Close()
	defer fmt.Println("off")

	for {

		data, err := bufio.NewReader(conn).ReadString('\x00')
		data = strings.TrimSpace(data)
		data = strings.TrimRight(data,"\x00")

		if err != nil {
			if err == io.EOF{
				return
			} else {
				fmt.Println("Error: " + err.Error())
				return
			}
		}

		fmt.Println("IN: " + data)

		outPacket := XMLPackets[data](conn, data) + "\u0000"
		conn.Write([]byte(outPacket))
		fmt.Println("OUT: " + outPacket)

	}
}

func policyFileRequest(conn net.Conn, packet string) string {
	return "<cross-domain-policy><allow-access-from domain='*' to-ports='" + conn.LocalAddr().String()[len(conn.LocalAddr().String()) - 4:] + "'/></cross-domain-policy>"
}

func versionCheck(conn net.Conn, packet string) string {
	return "<msg t='sys'><body action='apiOK' r='0'></body></msg>"
}

func randomKey(conn net.Conn, packet string) string {
	return "<msg t='sys'><body action='rndK' r='-1'><k>" + "" + "</k></body></msg>"
}

func login(conn net.Conn, packet string) string {
	return ""
}