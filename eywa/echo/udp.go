package echo

import (
	"eywa/utils"
	"log"
	"net"
)

const (
	BuffSize = 16 * 1024 * 1024
)

type UDP struct {
	Port string
	Conn *net.UDPConn
}

func (s UDP) Start(port string) {
	var err error
	_, s.Conn, err = utils.UDPServer(":"+port, BuffSize)
	if err != nil {
		log.Fatal(err)
	}

	go s.process()

	select {}
}

func (s *UDP) process() {
	buf := make([]byte, BuffSize)
	for {
		n, addr, err := s.Conn.ReadFrom(buf[0:])
		if err != nil {
			continue
		}
		s.Send(addr, buf[0:n])
	}
}

func (s *UDP) Send(addr net.Addr, msg []byte) bool {
	if _, err := s.Conn.WriteTo(msg, addr); err != nil {
		return false
	}

	return true
}
