// Copyright 2019-2021 eywa authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
package utils

import (
	"fmt"
	"net"
)

// TCPServer init and return an TCP address
func TCPServer(addr string) (net.Conn, net.Listener, error) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, nil, err
	}

	conn, err := ln.Accept()
	if err != nil {
		return nil, nil, err
	}
	return conn, ln, nil
}

// TCPClient create a TCP client
func TCPClient(addr string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// TCPSend send message
func TCPSend(conn net.Conn, message string) {
	fmt.Fprintf(conn, message+"\n")
}
