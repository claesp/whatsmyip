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
			a := []byte("HTTP/1.0 200 OK\nConnection: close\nContent-Type: text\n\n")
			t := conn.RemoteAddr().String()
			i := 0
			v6 := false
			if t[0] == '[' {
				v6 = true
				i = 1
			}

			for j := i; j < len(t); j++ {
				ch := t[j]
				if !v6 {
					if ch == ':' {
						break
					}
				} else {
					if ch == ']' {
						break
					}
				}

				a = append(a, byte(ch))
			}
			a = append(a, byte('\n'))

			conn.Write(a)
			conn.Close()
		}(c)
	}
}
