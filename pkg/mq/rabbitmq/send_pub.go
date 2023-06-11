package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func errorWithMsg(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

// 发布订阅模式: 生产者 => 交换机(fanout) => 队列 =》 消费者，一对多，多个消费者可以同时收到消息
func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	errorWithMsg(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	errorWithMsg(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	errorWithMsg(err, "Failed to declare an exchange")

	body := from(os.Args)
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	errorWithMsg(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

func from(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
