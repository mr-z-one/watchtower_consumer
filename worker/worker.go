package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	custom_color "watchtower_consumer/custom_Color"
	"watchtower_consumer/rabbitmq"
	AppHandler "watchtower_consumer/tools/handler"
	"watchtower_consumer/utils"
)

var number_Of_Worker int = runtime.NumCPU() - 2

func Worker(c context.Context, id int) {
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
				data, err := e.Execute(m.Args...)
				if utils.FailOnError(err, "error form Running Command  id : "+m.UUID, nil) {
					d.Nack(false, true)
					continue
				}
				custom_color.Succeed()("\n%s\n", string(data))

				d.Ack(false)
			}
		}
	}
}

func InitWorker(c context.Context) {
	for i := 0; i < number_Of_Worker; i++ {
		go Worker(c, i)
	}
}
