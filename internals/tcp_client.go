package pipie

import (
	"bufio"
	"encoding/json"
	"github.com/boltdb/bolt"
	"log"
	"net"
	"strconv"
)

type OnMessageFunc func(string)

type MqClient struct {
	Hostname string
	HostPort int
	DB       *bolt.DB
}

func (m MqClient) Receive(process OnMessageFunc) {

	// Ensure all routines finish before returning

	var count = 0
	conn, err := net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))

	if err != nil {
		log.Println(err)
	}
	for {
		count++

		if conn != nil {
			r := bufio.NewReader(conn)

			unserialized_message, err := r.ReadString('\n')

			if err != nil {
				conn, err = net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))
			} else if unserialized_message != "" {

				var message = Message{}

				b := []byte(unserialized_message)
				err := json.Unmarshal(b, &message)

				ack := Message{Key: message.Key, Value: "ack"}

				d, err := json.Marshal(ack)

				if err != nil {
					log.Print(err)
				}

				conn.Write([]byte(string(d) + "\n"))

				process(message.Value)
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

func (m MqClient) ReceiveAndSendAck(mq MqServer, process OnMessageFunc) {

	// Ensure all routines finish before returning

	var count = 0
	conn, err := net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))

	if err != nil {
		log.Println(err)
	}
	for {
		count++

		if conn != nil {
			r := bufio.NewReader(conn)

			unserialized_message, err := r.ReadString('\n')

			if err != nil {
				conn, err = net.Dial("tcp", m.Hostname+":"+strconv.Itoa(m.HostPort))
			} else if unserialized_message != "" {

				var message = Message{}

				b := []byte(unserialized_message)
				err := json.Unmarshal(b, &message)

				if err != nil {
					log.Print(err)
				}

				process(message.Value)

				log.Println("Sending ack ", message.Key)

				go func(key string) {
					mq.Send(key)
				}(message.Key)

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

func (m MqClient) ReceiveFromEternity(mq MqServer, process OnMessageFunc) {

	m.ReceiveAndSendAck(mq, process)
}
