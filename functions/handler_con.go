package functions

import (
	"fmt"
	"net"
	"strings"
	"sync"
)

type Info struct {
	Clients          map[net.Conn]string
	mut              sync.Mutex
	FilePath         string
	ClientAndMessage []string
}

func (a *Info) HandlerCon(con net.Conn) {
	defer func() {
		a.mut.Lock()
		name := a.Clients[con]
		delete(a.Clients, con)
		a.mut.Unlock()

		a.Disconnected(con, name)
		con.Close()
	}()

	buffer := make([]byte, 1024)

	_, err := con.Write([]byte("Enter your name: "))
	if err != nil {
		fmt.Println("Error sending prompt:", err)
		return
	}

	n, err := con.Read(buffer)
	if err != nil {
		return
	}

	name := strings.TrimSpace(string(buffer[:n]))
	help := false
	ret := isValidName(name)
	if ret == "" {
		help = true

		a.mut.Lock()
		a.Clients[con] = name
		a.mut.Unlock()
	}
	for !help {

		_, err := con.Write([]byte(ret))
		if err != nil {
			fmt.Println("Error sending invalid name message:", err)
			return
		}

		n, err := con.Read(buffer)
		if err != nil {
			return
		}

		name = strings.TrimSpace(string(buffer[:n]))
		ret = isValidName(name)
		if ret == "" {
			help = true
			a.mut.Lock()
			a.Clients[con] = name
			a.mut.Unlock()
		}

	}

	a.Chat(con)
}

func isValidName(name string) string {
	if len(name) < 3 {
		return "Name is too short. Must be at least 3 characters.\n"
	} else if len(name) > 25 {
		return "Name is too long. Maximum 25 characters allowed.\n"
	}
	return ""
}

func (a *Info) Chat(con net.Conn) {
	a.mut.Lock()
	if len(a.ClientAndMessage) > 0 {
		for _, msg := range a.ClientAndMessage {
			_, err := con.Write([]byte(msg + "\n"))
			if err != nil {
				fmt.Println("Error sending old messages:", err)
				a.mut.Unlock()
				return
			}
		}
	}
	a.mut.Unlock()

	buffer := make([]byte, 1024)
	for {
		name := a.Clients[con]
		n, err := con.Read(buffer)
		if err != nil {
			fmt.Println(name,"disconnected!!	")
			return
		}

		message := strings.TrimSpace(string(buffer[:n]))

		NameAndMessage := "[" + name + "]: " + message

		a.mut.Lock()
		a.ClientAndMessage = append(a.ClientAndMessage, NameAndMessage)
		a.mut.Unlock()

		a.mut.Lock()
		for client := range a.Clients {
			if client != con {
				_, err := client.Write([]byte(NameAndMessage + "\n"))
				if err != nil {
					fmt.Println("Error sending message to client:", err)
				}
			}
		}
		a.mut.Unlock()
	}
}
func (a *Info) Disconnected(con net.Conn, name string) {
	a.mut.Lock()
	defer a.mut.Unlock()

	disconnectMessage := fmt.Sprintf("[%s] has disconnected\n", name)
	for client := range a.Clients {
		if client != con {
			_, err := client.Write([]byte(disconnectMessage))
			if err != nil {
				fmt.Println("Error sending disconnection message:", err)
			}
		}
	}
}	
