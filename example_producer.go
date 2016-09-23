package main

import (
	"time"
	"github.com/debarshri/pipie/internals"
)

func main(){
	mq := pipie.ServerStreamAtPort(8081)

	for{
		mq.Send(time.Now().String())
		time.Sleep(100*time.Millisecond)
	}

	mq.Stop()

}