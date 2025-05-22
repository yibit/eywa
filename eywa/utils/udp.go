// Copyright 2019-2021 eywa authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.
package utils

import (
	"log"
	"net"
)

type UDPNode struct {
	Conn     *net.UDPConn
	Addr     net.Addr
	MetaAddr string
}

func NewUDPNode(address string) *UDPNode {
	return &UDPNode{
		MetaAddr: address,
	}
}

// UDPServer init and return an UDP address
func UDPServer(address string, buffSize int) (*net.UDPAddr, *net.UDPConn, error) {
	addr := resolveUDPAddr(address)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return nil, nil, err
	}

	conn.SetReadBuffer(buffSize)

	return addr, conn, nil
}

// UDPClient init and return an UDP client
func UDPClient(address string) (*net.UDPAddr, *net.UDPConn, error) {
	addr := resolveUDPAddr(address)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return nil, nil, err
	}

	return addr, conn, nil
}

// resolveUDPAddr resolve UDP address
func resolveUDPAddr(address string) *net.UDPAddr {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(nil, err)
	}

	return addr
}

func Send(conn *net.UDPConn, message []byte) (int, error) {
	return conn.Write(message)
}

func (m *UDPNode) InitClient(address string) (*UDPNode, error) {
	var err error
	m.Addr, m.Conn, err = UDPClient(address)
	if err != nil {
		return nil, err
	}

	m.MetaAddr = address

	return m, nil
}

func (m *UDPNode) InitServer(address string, buffSize int) (*UDPNode, error) {
	var err error
	m.Addr, m.Conn, err = UDPServer(address, buffSize)
	if err != nil {
		return nil, err
	}

	m.MetaAddr = address

	return m, nil
}
