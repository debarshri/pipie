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
	"encoding/json"
	"sync"
)

var wg sync.WaitGroup

func Start() MqServer {
	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8080")

	// accept connection on port

	return MqServer{listener: ln}

}

func ServerStreamAtPort(port int) MqServer {
	fmt.Println("Starting producer...", port)

	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(port))

	db, err := bolt.Open("pipie.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	mq := MqServer{listener: ln, DB:db}

	//go mq.startService(strconv.Itoa(port + 1))

	return mq
}

func ServerStreamAtPortWithDBLoc(port int, dblocation string) MqServer {
	fmt.Println("Starting producer...", port)

	ln, _ := net.Listen("tcp", ":"+strconv.Itoa(port))

	db, err := bolt.Open(dblocation, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	mq := MqServer{listener: ln, DB:db}

	//go mq.startService(strconv.Itoa(port + 1))

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

	mq.DB.Update(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("data"))

		if c != nil {

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

	con , _:= m.listener.Accept()

	message := Message{Key:key, Value:data}

	serialzed_message, err := json.Marshal(message)

	log.Print("Sending key ", key)

	_, err = con.Write([]byte(string(serialzed_message) + "\n"))

	if err != nil {
		log.Println(err)
	}

	con.Close()
}

func (m MqServer) Send(data string) {

	key := strconv.FormatInt(time.Now().UnixNano(), 10)
	m.send(key,data)
}

func (m MqServer) DeleteKey(data string) {

	go m.DB.Update(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("data"))

		err := c.Delete([]byte(data))

		if err != nil {
			log.Println(err)
		}

		log.Println("Deleting key", data)

		return nil
	})
}

func (m MqServer) Persist(payload string) {

	key := strconv.FormatInt(time.Now().UnixNano(), 10)
	m.DB.Update(func(tx *bolt.Tx) error {

		c, _ := tx.CreateBucketIfNotExists([]byte("data"))

		log.Println("Persisting data with key", key)
		c.Put([]byte(key), []byte(payload))

		return nil
	})
}

func (m MqServer) Flush() {

	go m.DB.Update(func(tx *bolt.Tx) error {

		c := tx.Bucket([]byte("data"))

		if c != nil {

			cursor := c.Cursor()
			var count int = 0

			for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
				log.Println("Working for the key ", string(k))

				count = count+1
				go m.send(string(k),string(v))
			}

			log.Println("Total messages in the queue ",count)
		}

		return nil
	})
}



