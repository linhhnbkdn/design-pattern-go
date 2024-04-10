package main

import (
	"log"
)

type Content struct {
	url string
	body string
	data any // can be anything
}


type Handler func(c *Content) error

func Welcome(c *Content) error {
	c.body = "Welcome"
	log.Println("- A simple web drawler", c.url)
	return nil
}

func CheckUrl(c *Content) error {
	log.Println("- Checking the url", c.url)
	return nil
}

func FetchHtml(c *Content) error {
	log.Println("- Fetching the html from url", c.data)
	return nil
}

func ParseHtml(c *Content) error {
	log.Println("- Parsing the html")
	c.data = map[string]string{
		"title": "Google",
		"content": "Google is a search engine",
	}
	return nil
}

type HandlerNode struct {
	hdl Handler
	next *HandlerNode
}

func (hn *HandlerNode) Handle(url string) error {
	ctx := Content{url: url}

	if hn == nil || hn.hdl == nil {
		return nil
	}
	if err := hn.hdl(&ctx); err != nil {
		return err
	}
	return hn.next.Handle(url)
}

func NewCrawler(handlers ...Handler) HandlerNode {
	node := HandlerNode{}
	if len(handlers) > 0 {
		node.hdl = handlers[0]
	}
	current_node := &node
	for i := 1; i < len(handlers); i++ {
		current_node.next = &HandlerNode{hdl: handlers[i]}
		current_node = current_node.next
	}

	return node
}

type WebDrawler struct {
	handlers HandlerNode
}

func (wc WebDrawler) Draw(url string) {
	if err := wc.handlers.Handle(url); err != nil {
		log.Println("Error:", err)
	}
}

func main(){
	WebDrawler{
		handlers: NewCrawler(
			Welcome,
			CheckUrl,
			FetchHtml,
			ParseHtml,
		),
	}.Draw("https://www.google.com")
}


