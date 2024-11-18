package main

import (
	"fmt"
	"net"
	"os"

	"nc/functions"
)

func main() {
	port := "8989"
	if len(os.Args) == 2 && isValid(os.Args[1]) {
		port = os.Args[1]
	}
	ls, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	fmt.Println("Server listening on port", port)

	serverInfo := &functions.Info{
		Clients: make(map[net.Conn]string),
	}
	for {
		con, err := ls.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go serverInfo.HandlerCon(con)
	}
}

func isValid(arg string) bool {
	for _, char := range arg {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
