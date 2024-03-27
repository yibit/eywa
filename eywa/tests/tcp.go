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

	conn, err := utils.TCPClient(":" + Port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("the TCP server is %s\n", conn.RemoteAddr().String())
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)

		if strings.ToUpper(strings.TrimSpace(text)) == "QUIT" {
			fmt.Println("> exiting...")
			return
		}

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Printf("<< %s", message)
	}
}
