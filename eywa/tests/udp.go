package main

import (
	"bufio"
	"eywa/utils"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	Port = "1037"
)

func main() {
	flag.StringVar(&Port, "p", Port, "`port` of server")
	flag.Parse()

	_, conn, err := utils.UDPClient(":" + Port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("the UDP server is %s\n", conn.RemoteAddr().String())
	defer conn.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.ToUpper(strings.TrimSpace(text)) == "QUIT" {
			fmt.Println("> exiting...")
			return
		}
		buffer := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("<< %s", string(buffer[0:n]))
	}
}
