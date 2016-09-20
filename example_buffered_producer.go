package main

import (
	"time"
	"github.com/debarshri/pipie/internals"
)

func main(){
	mq := pipie.ServerStreamAtPort(8081)

	for{
		mq.BufferedSend(time.Now().String())
	}

	mq.Stop()

}