package pipie

func CreateConsumer(node Node) Consumer {

	var dbLocation string
	if node.ProducerHostName != "" {

		if node.DBLocation == "" {
			dbLocation = "pipie-consumer.db"
		} else {
			dbLocation = node.DBLocation
		}

		server := ServerStreamAtPortWithDBLoc(node.ConsumerPort, dbLocation)
		return Consumer{
			Client: MqClient{Hostname: node.ProducerHostName,
				HostPort: node.ProducerPort,
				DB:       server.DB,
			},
			Server: server,
		}
	}

	return Consumer{}
}

func (c Consumer) Receive(process OnMessageFunc) {
	c.Client.ReceiveAndSendAck(c.Server, process)
}
