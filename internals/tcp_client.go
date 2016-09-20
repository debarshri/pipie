package pipie

import (
	"bufio"
	"log"
	"net"
	"net/http"
	"strconv"
)

type OnMessageFunc func(string)

type MqClient struct {
	Hostname string
	HostPort int
}

func (m MqClient) Receive(process OnMessageFunc) {

	var count = 0
	conn, err := net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))

	if err != nil {
		log.Println(err)
	}
	for {
		count++

		if conn != nil {
			r := bufio.NewReader(conn)

			message, err := r.ReadString('\n')

			if err != nil {
				conn, err = net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))
			} else if message != "" {
				process(message)
			}
		}

		if count%1000000000 == 0 {
			log.Println("Recaliberate the connection")
			conn, err = net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))

			if err != nil {
				log.Println(err)
			}
		}
	}
}

func (m MqClient) ReceiveFromEternity(process OnMessageFunc) {

	go http.Get("http://"+m.Hostname+":"+strconv.Itoa(m.HostPort+1)+"/all")
	m.Receive(process)

}
