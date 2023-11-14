package main

import "fmt"

type ChatMessage struct {
	Content string
	SenderName string
	SenderAvatar []byte
}

func main()	{
	chatMessages := []ChatMessage {
		{
			Content: "Hello",
			SenderName: "John",
			SenderAvatar: make([]byte, 1024*300),
		},
		{
			Content: "Hi",
			SenderName: "Jane",
			SenderAvatar: make([]byte, 1024*300),
		},
		{
			Content: "How are you?",
			SenderName: "John",
			SenderAvatar: make([]byte, 1024*300),
		},
		{
			Content: "I'm fine",
			SenderName: "Jane",
			SenderAvatar: make([]byte, 1024*300),
		},
	}
	fmt.Println(chatMessages)
}

// We increase the memory due to SenderAvatar and SenderName create new memory
// --> Just only create one time and use it for all chatMessages


