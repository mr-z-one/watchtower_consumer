package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	data "watchtower_consumer/Data"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/rabbitmq"
	AppHandler "watchtower_consumer/tools/handler"
	"watchtower_consumer/utils"
)

var number_Of_Worker int = runtime.NumCPU() - 2

func Worker(c context.Context, data_channel chan<- *data.MessageData) {
	msg, ch := rabbitmq.GetMessage()
	defer ch.Close()

	for {
		select {
		case <-c.Done():
			fmt.Println("Done")

			return

		default:

			for d := range msg {
				m := rabbitmq.Message{}
				json.Unmarshal(d.Body, &m)

				e, err := AppHandler.Get_Program(m.CMD)

				if utils.FailOnError(err, "command Not Found with id : "+m.UUID, nil) {
					d.Nack(false, false)
					continue
				}
				res, err := e.Execute(m.Args...)
				if utils.FailOnError(err, "error form Running Command  id : "+m.UUID, nil) {
					d.Nack(false, true)
					continue
				}
				custom_color.Succeed()("\n%s\n", "cmd are finish successful")
				messageData := data.NewMessageData(m.UUID, res...)

				data_channel <- messageData
				d.Ack(false)
			}
		}
	}
}
func HandleData(data_channel <-chan *data.MessageData) {
	for message := range data_channel {
		custom_color.Succeed()("msg_uuid %s \n", message.Msg_UUID)
		custom_color.Succeed()("\tmsg_response %s \n", string(message.Response))
	}
}
func InitWorker(c context.Context, data chan<- *data.MessageData) {
	for i := 0; i < number_Of_Worker; i++ {
		go Worker(c, data)
	}
}
