#### Pipie

Pure golang, embedded point to point messaging library

### Getting started. 

It is as basic as it can get. Simple to use. Simple to extend.

Conceptually, this messaging library has a producer and a consumer.
Producer looks like this

```
import (
	"github.com/debarshri/pipie/broker"
)

func main(){
	mq := pipie.StartWithPort("8081")
	//mq := pipie.Start() Starts at 8080
	mq.Publish("Some data)
}
```

Consumer looks like this

```
import (
	"fmt"
	"github.com/debarshri/pipie/broker"
)

func main() {

	q := pipie.MqClient{Hostname:"localhost:8081"}

	q.Receive(func(data string){
		fmt.Println(data)
	})
}
```

As you can see you can pass the func where you receive data, process the data in that function.

You can have multiple consumers, in that scenario the messages are published in round robin fashion.

It doesn't support message broadcasting. It doesn't support persistance of undelivered messages, acknowledgments. 
Goddamit, theres so much work to be done.

This library was written for IoT kind of setting, just a side note, if you like side note. I like sidenotes.

P.S. It would be really nice if someone would like to mentor me in the project. 
If you read the code, you would know why I need a mentor.