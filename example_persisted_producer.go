package main

import (
	"time"
	"github.com/debarshri/pipie/internals"
)

func main(){
	mq := pipie.ServerStreamAtPort(8081)

	for{
		mq.PersistedSend(time.Now().String())
		time.Sleep(1*time.Second)
	}

	mq.Stop()

}