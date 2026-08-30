package rabbitmq

import "github.com/google/uuid"

var (
	SUBFINDER string = "subfinder"
)

type Message struct {
	UUID string   `json:"uuid" bson:"uuid"`
	CMD  string   `json:"cmd" bson:"cmd"`
	Args []string `json:"args" bson:"args"`
}

func NewMessage(cmd string, args ...string) *Message {
	return &Message{UUID: uuid.NewString(), CMD: cmd, Args: args}

}
