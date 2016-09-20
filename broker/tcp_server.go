package pipie

import (
	"fmt"
	"net"
	"log"
)

type Send func() string


type MqServer struct {
    listener net.Listener
    ttl int64
}

func Start() MqServer {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port

	return MqServer{listener: ln}

}

func StartWithPort(port string) MqServer {
	fmt.Println("Starting producer...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":"+port)

	return MqServer{listener: ln}

}

func (m MqServer) Publish(data string){

	con, _ := m.listener.Accept()

	_, err := con.Write([]byte(data + "\n"))

	if err != nil{
		log.Println(err)
	}

	con.Close()

}
