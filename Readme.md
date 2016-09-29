#### Pipie

Pure golang, embedded point to point messaging library

### Getting started. 

It is as basic as it can get. Simple to use. Simple to extend.

Producer starts two services, publish service and subscribe service

When a producer publishes, the message is put into a bufferpool.

Bufferpool is flush every time period t.

On receiving an acknowledgment the message is removed from bufferpool. In pipie, bufferpool is persisted on disk.

On consumer side, when a message is received, pipie consumer first persists it and sends our a acknowledgement on subscribe service.
Application can then pick up message from this persisted bufferpool in time t. This is how, the backpressure is handled in pipie.

You can pass the func where you receive data, process the data in that function.

You can have multiple consumers, in that scenario the messages are published in round robin fashion.

It doesn't support message broadcasting yet.

This library was written for IoT kind of setting, just a side note, if you like side note. I like sidenotes.

P.S. It would be really nice if someone would like to mentor me in the project. 
If you read the code, you would know why I need a mentor.