package main

import "fmt"

type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.send(message)
}

type Notifier interface {
	send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) send(message string) {
	fmt.Printf("sending email: %s\n", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) send(message string) {
	fmt.Printf("sending SMS: %s\n", message)
}

type SlackNotifier struct{}

func (s *SlackNotifier) send(message string) {
	fmt.Printf("sending Slack message: %s\n", message)
}

func CreateNotifier(notificationStype string) Notifier {
	if notificationStype == "email" {
		return &EmailNotifier{}
	} else if notificationStype == "sms" {
		return &SMSNotifier{}
	} else {
		return &SlackNotifier{}
	}
}
func main() {
	notifier := CreateNotifier("email")
	s := NotificationService{notifier: notifier}
	s.SendNotification("Hello world!")

	notifier = CreateNotifier("sms")
	s = NotificationService{notifier: notifier}
	s.SendNotification("Hello world!")

	notifier = CreateNotifier("slack")
	s = NotificationService{notifier: notifier}
	s.SendNotification("Hello world!")
}
