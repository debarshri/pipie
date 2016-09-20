package pipie

import (
	"bufio"
	"net"
	"time"
	"log"
)

type OnMessageFunc func(string)

type MqClient struct {
	Hostname string
}

func (m MqClient) Receive(process OnMessageFunc) {

	var count = 0
	conn, err := net.Dial("tcp", m.Hostname)

	if err != nil {
		log.Println(err)
	}
	for {
		count++

		if conn != nil {
			r := bufio.NewReader(conn)

			message, err := r.ReadString('\n')

			if err != nil {
				conn, _ = net.Dial("tcp", m.Hostname)
			} else if message != "" {
				process(message)
			}
		}

		if count%10000000 == 0 {

			log.Println("Recaliberate the connection")
			conn, err = net.Dial("tcp", m.Hostname)

			if err != nil {
				log.Println(err)
				time.Sleep(1000*time.Millisecond)
			}
		}
	}

}


