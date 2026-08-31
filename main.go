package main

import (
	"context"
	"os/signal"
	"runtime"
	"syscall"
	"time"
	data "watchtower_consumer/Data"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/env"
	"watchtower_consumer/tools"
	AppHandler "watchtower_consumer/tools/handler"
	"watchtower_consumer/worker"
)

var BUFFER_CAPACITY int = runtime.NumCPU() * 2

func main() {
	var data = make(chan *data.MessageData, BUFFER_CAPACITY)

	env.Init()
	AppHandler.Register_All()

	if tools.Check_Proxy() {
		custom_color.Succeed()("[+] Proxy Turn on\n")
	} else {
		custom_color.Warning()("[!] Proxy Turn off\n")
	}

	// msg, ch := rabbitmq.GetMessage()
	// defer ch.Close()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	worker.InitWorker(ctx, data)
	go worker.HandleData(data)
	<-ctx.Done()
	defer stop()

	time.Sleep(1 * time.Second)
	//stop()

	// e, _ := AppHandler.Get_Program(Constants.APP_SUBFINDER)
	// e.Execute("sess.sku.ac.ir")

	// go func(msg <-chan amqp091.Delivery) {
	// 	for d := range msg {
	// 		custom_color.Succeed()("message relives {%s}\n", string(d.Body))
	// 		d.Nack(false, true)
	// 	}
	// }(msg)

	//select {}
}
