package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Text to send: ")

		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(conn, text+"\n")

		if strings.Trim(text, "\n") == "exit" {
			os.Exit(0)
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Message from server: " + message)
	}
}
