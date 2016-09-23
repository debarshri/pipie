package pipie

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/go-martini/martini"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

type MqServer struct {
	listener  net.Listener
	queuesize int
	ttl       int64
}

type Message struct {
	Key string
	Value string
}

var max_poolsize int = 10
var curr_poolsize int = 0
var queue []string = make([]string, 0)

func Start() MqServer {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port

	return MqServer{listener: ln}

}

func ServerStreamAtPort(port int) MqServer {
	fmt.Println("Starting producer...")

	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(port))

	mq := MqServer{listener: ln}
	go mq.startService(strconv.Itoa(port + 1))

	return mq

}

func (mq MqServer) startService(port string) {

	m := martini.Classic()

	os.Setenv("PORT", port)

	m.Get("/", func() string {
		return "MQ service is on"
	})

	m.Get("/all", func() string {

		mq.update()

		return "Done"
	})

	m.Run()
}

func (mq MqServer) update() {
	db, err := bolt.Open("pipie.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.View(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("data"))

		if c != nil {
			if err != nil {
				log.Fatal("Doesnt work")
			}
			cursor := c.Cursor()

			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				log.Println("Working for the key ", string(k))
				mq.Send(string(v))
				c.Delete(k)
			}
		}

		return nil
	})
}

func (m MqServer) send(key string, data string) {
	con, _ := m.listener.Accept()

	_, err := con.Write([]byte(data + "\n"))

	if err != nil {
		log.Println(err)
	}

	con.Close()
}

func (m MqServer) Send(data string) {
	key := strconv.FormatInt(time.Now().UnixNano(), 10)

	m.send(key,data)
}

//func (m MqServer) BufferedSend(data string) {
//
//	log.Println("Send size ", len(queue))
//
//	if len(queue) < 100000{
//		queue = append(queue, data)
//	} else {
//		queue = queue[1:]
//		queue = append(queue, data)
//	}
//
//	curr_poolsize++
//
//	if curr_poolsize < max_poolsize {
//		go m.Send(data)
//		curr_poolsize = curr_poolsize - 1
//	}
//}

func (m MqServer) PersistedSend(data string) {

	db, err := bolt.Open("pipie.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	key := strconv.FormatInt(time.Now().UnixNano(), 10)

	err = db.Update(func(tx *bolt.Tx) error {
		c, _ := tx.CreateBucketIfNotExists([]byte("data"))

		log.Println("Adding data")
		c.Put([]byte(key), []byte(data))
		return nil
	})

	log.Println("Send size ", len(queue))

	curr_poolsize = curr_poolsize + 1

	go m.send(key, data)
}

func (m MqServer) Stop() {
	m.listener.Close()
}
