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
	for {
		buf := make([]byte, BuffSize)

		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("%s\n", err.Error())
			return
		}

		if _, err = conn.Write(buf[:n]); err != nil {
			log.Printf("%s\n", err.Error())
			continue
		}
	}
}
