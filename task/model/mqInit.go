package model

import (
	"github.com/rabbitmq/amqp091-go"
)

var MQ *amqp091.Connection

func RabbitMQ(connString string)  {
	conn,err := amqp091.Dial(connString)
	if err != nil{
		panic(err)
	}
	MQ = conn

}
