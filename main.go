package main

import (
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/env"
	"watchtower_consumer/rabbitmq"
	"watchtower_consumer/tools"
	AppHandler "watchtower_consumer/tools/handler"

	"github.com/rabbitmq/amqp091-go"
)

func main() {

	env.Init()
	AppHandler.Register_All()

	if tools.Check_Proxy() {
		custom_color.Succeed()("[+] Proxy Turn on\n")
	} else {
		custom_color.Warning()("[!] Proxy Turn off\n")
	}

	msg, ch := rabbitmq.GetMessage()
	defer ch.Close()
	// e, _ := AppHandler.Get_Program(Constants.APP_SUBFINDER)
	// e.Execute("sess.sku.ac.ir")

	go func(msg <-chan amqp091.Delivery) {
		for d := range msg {
			custom_color.Succeed()("message relives {%s}\n", string(d.Body))
			d.Nack(false, true)
		}
	}(msg)

	select {}
}
