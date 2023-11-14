package main

import "fmt"

type Sender struct {
	Name string
	Avatar []byte
}

type ChatMessage struct {
	Content string
	Sender *Sender
}

type SenderFactory struct {
	senders map[string]*Sender
}

func (sf *SenderFactory) GetSender(name string) *Sender {
	return sf.senders[name]
}

func main() {
	JoseSender := Sender{
		Name: "Jose",
		Avatar: make([]byte, 1024*300),
	}
	JaneSender := Sender{
		Name: "Jane",
		Avatar: make([]byte, 1024*300),
	}

	senderFactory := SenderFactory{
		senders: map[string]*Sender{
			"Jose": &JoseSender,
			"Jane": &JaneSender,
		},
	}

	fmt.Println(len(senderFactory.senders))
	chatConversation := []ChatMessage{
		{
			Content: "Hello",
			Sender: senderFactory.GetSender("Jose"),
		},
		{
			Content: "Hi",
			Sender: senderFactory.GetSender("Jane"),
		},
	}
	fmt.Println(chatConversation)

}
