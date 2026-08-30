package rabbitmq

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"watchtower_consumer/utils"

	"github.com/rabbitmq/amqp091-go"
)

var connection *amqp091.Connection

func Connect() *amqp091.Connection {
	if connection != nil {
		return connection
	}
	conn, err := amqp091.Dial("amqp://guest:guest@localhost:5672/")
	utils.FailOnErrorPanic(err, "", nil)

	return conn
}

func CreateChannel(conn *amqp091.Connection) *amqp091.Channel {

	ch, err := conn.Channel()
	utils.FailOnErrorPanic(err, "", nil)

	return ch
}
func SendMessage(message *Message, callback func(ack bool, message *Message)) {
	conn := Connect()
	if conn == nil {
		panic("rabbitmq connection Lost")
	}
	ch := CreateChannel(conn)
	defer ch.Close()
	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durability
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	utils.FailOnErrorPanic(err, "", nil)
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnErrorPanic(err, "", nil)

	ch.Confirm(false)

	parentCtx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,    // Ctrl+C
		syscall.SIGTERM, // Docker stop, Kubernetes termination
	)
	defer stop()
	const confirmTimeout = 5 * time.Second
	jsonData, _ := json.Marshal(message)

	confirm, err := ch.PublishWithDeferredConfirmWithContext(parentCtx, "", q.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        jsonData,
	})

	defer stop()
	ctx, cancel := context.WithTimeout(parentCtx, confirmTimeout)
	acked, err := confirm.WaitContext(ctx)
	defer cancel()

	if err != nil {
		log.Printf("Confirmation wait failed: %v", err)
		callback(false, message)
		return
	}
	if !acked {
		log.Printf("Message %d was nacked by the broker", confirm.DeliveryTag)
		callback(false, message)
		return
	}
	callback(true, message)

}

func GetMessage() (_ <-chan amqp091.Delivery, channel *amqp091.Channel) {
	conn := Connect()
	if conn == nil {
		panic("rabbitmq connection Lost")
	}
	ch := CreateChannel(conn)

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durability
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		amqp091.Table{
			amqp091.QueueTypeArg: amqp091.QueueTypeQuorum,
		},
	)
	utils.FailOnErrorPanic(err, "", nil)
	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	utils.FailOnErrorPanic(err, "", nil)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnErrorPanic(err, "Failed to register a consumer", nil)

	return msgs, ch
}
