package main

import "fmt"

type Notifier interface {
	SendMessage(msg string)
}

type EmailNotifier struct {
}

func (EmailNotifier) SendMessage(msg string) {
	// Send email
	fmt.Printf("Send email: %s Sender: Email \n", msg)
}

type SMSNotifier struct {
}

func (SMSNotifier) SendMessage(msg string) {
	// Send SMS
	fmt.Printf("Send SMS: %s Sender: SMS \n", msg)
}


type TelegramNotifier struct {
}

func (TelegramNotifier) SendMessage(msg string) {
	// Send Telegram
	fmt.Printf("Send Telegram: %s Sender: Telegram \n", msg)
}

type SlackNotifier struct {
}

func (SlackNotifier) SendMessage(msg string) {
	// Send Slack
	fmt.Printf("Send Slack: %s Sender: Slack \n", msg)
}

type StackNotifier struct {
	core *StackNotifier
	notifier Notifier
}

func (s StackNotifier) SendMessage(msg string) {
	s.notifier.SendMessage(msg)
	if s.core != nil {
		s.core.SendMessage(msg)
	}
}

type NotifierService struct {
	notifier Notifier
}

func (n NotifierService) SendMessage(msg string) {
	n.notifier.SendMessage(msg)
}

type NotifierDecorator struct {
	core *NotifierDecorator
	notifier Notifier
}

func (n NotifierDecorator) SendMessage(msg string) {
	n.notifier.SendMessage(msg)
	if n.core != nil {
		n.core.SendMessage(msg)
	}
}

func (n NotifierDecorator) AddDecorator(notifier Notifier) NotifierDecorator {
	return NotifierDecorator{
		core:     &n,
		notifier: notifier,
	}
}

func main() {
	// notifier := NotifierService{
	// 	notifier: EmailNotifier{},
	// }
	// notifier.SendMessage("Hello")
	notifierMaster := NotifierDecorator{
		notifier: EmailNotifier{},
	}.AddDecorator(SMSNotifier{}).AddDecorator(TelegramNotifier{}).AddDecorator(SlackNotifier{})
	notifierMaster.SendMessage("Hello")
}


