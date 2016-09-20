package main

import (
	"time"
	"github.com/debarshri/pipie/internals"
)

func main(){
	mq := pipie.StartWithPort("8081")

	for{
		mq.Publish(time.Now().String())
		time.Sleep(100*time.Millisecond)
	}

}