package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		message = strings.Trim(message, "\n")

		if strings.Trim(message, "\n") == "exit" {
			os.Exit(0)
		}

		var NewMessage string

		number, err := strconv.Atoi(message)

		if err != nil {
			NewMessage = strings.ToUpper(message)
		} else {
			NewMessage = strconv.Itoa(number * 2)
		}

		conn.Write([]byte(NewMessage + "\n"))
	}
}