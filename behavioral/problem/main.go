package main

import "fmt"

type NotificationService struct {
	notificationType string
}

func (n *NotificationService) SendNotification(message string) {
	if n.notificationType == "email" {
		fmt.Printf("Sending email: %s\n", message)
	} else if n.notificationType == "sms" {
		fmt.Printf("Sending SMS: %s\n", message)
	} else {
		fmt.Printf("Sending message: %s\n", message)
	}
}

func main() {
	s := NotificationService{notificationType: "email"}
	s.SendNotification("Hello world!")
}
