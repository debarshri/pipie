package main

import (
	"fmt"
	"github.com/debarshri/pipie/internals"
)

func main() {

	q := pipie.MqClient{Hostname:"localhost", HostPort:8081}

	q.ReceiveFromEternity(func(data string){
		fmt.Println(data)
	})

}
