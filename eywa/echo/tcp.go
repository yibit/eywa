package echo

import (
	"bufio"
	"eywa/utils"
	"log"
	"net"
)

type TCP struct {
	Port     string
	Conn     net.Conn
	Listener net.Listener
}

func (s *TCP) Start(port string) {
	var err error
	s.Conn, _, err = utils.TCPServer(":" + port)
	if err != nil {
		log.Fatal(err)
	}

	go s.process()
}

func (s *TCP) process() {
	for {
		message, err := bufio.NewReader(s.Conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		s.Conn.Write([]byte(message + "\n"))
	}
}
