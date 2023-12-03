package main

import "fmt"

type NotificationService struct {
	notifier Notifier
}

func (n *NotificationService) SendNotification(message string) {
	n.notifier.Send(message)
}

type Notifier interface {
	Send(message string)
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) {
	fmt.Printf("Sending email: %s\n", message)
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) {
	fmt.Printf("Sending SMS: %s\n", message)
}

type SlackNotifier struct{}

func (s *SlackNotifier) Send(message string) {
	fmt.Printf("Sending Slack message: %s\n", message)
}

func main() {
	s := NotificationService{notifier: &EmailNotifier{}}
	s.SendNotification("Hello world!")

	s = NotificationService{notifier: &SMSNotifier{}}
	s.SendNotification("Hello world!")

	s = NotificationService{notifier: &SlackNotifier{}}
	s.SendNotification("Hello world!")
}
