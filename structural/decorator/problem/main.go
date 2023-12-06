package main

import (
	"fmt"
)

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

type EmailSMSNotifier struct {
	emailNotifier EmailNotifier
	smsNotifier  SMSNotifier
}

func (e EmailSMSNotifier) SendMessage(msg string) {
	e.emailNotifier.SendMessage(msg)
	e.smsNotifier.SendMessage(msg)
}

// Nếu muốn gửi cùng lúc nhiều thông báo thì sẽ phải tạo nhiều tổ hợp kiểu EmailSMSNotifier
// Vì vậy ta sẽ sử dụng decorator pattern để giải quyết vấn đề này

