package main

import "fmt"

type WebDrawler struct {
}

func (w *WebDrawler) Draw(url string) {
	fmt.Println("A simple web drawler", url)
	fmt.Println("- Checking the url")
	fmt.Println("- Fetching the html from url")
	fmt.Println("- Parsing the html")
	fmt.Println(" - Done")
}


func main(){
	var drawler WebDrawler
	drawler.Draw("https://www.google.com")
}
