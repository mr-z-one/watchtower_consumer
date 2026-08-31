package data

type MessageData struct {
	Msg_UUID string `json:"msg_uuid"`
	Response []byte `json:"response"`
}

func NewMessageData(uuid string, args ...byte) *MessageData {

	return &MessageData{Msg_UUID: uuid, Response: args}
}
