package client

import "net"

func Connect() {
	conn, err := net.Dial("tcp", "127.0.0.1:6969")

	if err != nil {
		panic(err)
	}

	_, err = conn.Write([]byte("bellend\u0000"))

	if err != nil {
		panic(err)
	}
}
