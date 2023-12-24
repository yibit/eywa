package echo

import (
	"bufio"
	"eywa/utils"
	"log"
	"net"
)

type TCP struct {
	Port     string
	Listener net.Listener
}

func (s TCP) Start(port string) {
	var err error
	s.Listener, err = utils.TCPServer(":" + port)
	if err != nil {
		log.Fatal(err)
	}

	defer s.Listener.Close()

	go s.process()

	select {}
}

func (s *TCP) process() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Printf("%s\n", err.Error())
			continue
		}

		go run(conn)
	}
}

func run(conn net.Conn) {
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Printf("%s\n", err.Error())
			return
		}
		conn.Write([]byte(message + "\n"))
	}
}
