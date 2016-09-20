package main

import (
	"fmt"
	"github.com/debarshri/pipie/internals"
)

func main() {

	q := pipie.MqClient{Hostname:"localhost:8081"}

	q.Receive(func(data string){
		fmt.Println(data)
	})
}
