package main

import (
	"net"
)

func main() {
	l, err := net.Listen("tcp4", "0.0.0.0:8081")
	if err != nil {
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			return
		}

		go func(conn net.Conn) {
			a := []byte("HTTP/1.1 200 OK\nConnection: close\nContent-Type: text\n\n")
			t := conn.RemoteAddr().String()

			for i := 0; i < len(t); i++ {
				if t[i] == ':' {
					break
				}

				a = append(a, byte(t[i]))
			}
			a = append(a, byte('\n'))

			conn.Write(a)
			conn.Close()
		}(c)
	}
}
