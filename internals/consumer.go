package pipie


func CreateConsumer(node Node)(Consumer){

	if node.ProducerHostName != "" {

		server := ServerStreamAtPortWithDBLoc(node.ConsumerPort, "pipie-consumer.db")
		return Consumer{
			Client:MqClient{Hostname:node.ProducerHostName,
				HostPort:node.ProducerPort,
				DB:server.DB,
			},
			Server:server,
		}
	}

	return Consumer{}
}

func (c Consumer) Receive(process OnMessageFunc){
	c.Client.ReceiveAndSendAck(c.Server, process)
}

