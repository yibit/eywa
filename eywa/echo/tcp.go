package echo

import (
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
	buf := make([]byte, BuffSize)
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Printf("%s\n", err.Error())
			break
		}

		if _, err = conn.Write(buf[:n]); err != nil {
			log.Printf("%s\n", err.Error())
			continue
		}
	}

	defer conn.Close()
}
