package main

import (
	"encoding/json"
	"mall-client/rabbitmq"
)

func main() {

	qe := rabbitmq.QueueAndExchange{
		QueueName:    "test_queueName",
		ExchangeName: "test_exchangeName",
		ExchangeType: "direct",
		RoutingKey:   "test_routingKey",
	}

	mq := rabbitmq.NewRabbitMq(qe)

	mq.ConnMq()
	defer mq.CloseMq()
	mq.OpenChan()
	defer mq.CloseChan()

	data := map[string]interface{}{
		"name": "test_name",
		"id":   "id",
	}

	body, _ := json.Marshal(data)

	mq.PublishMsg(string(body))

}
