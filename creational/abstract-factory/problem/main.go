package main

import "fmt"

type Food interface {
	Eat()
}

type Drink interface {
	Drink()
}

type Voucher struct {
	Drink
	Food
}

type Coffee struct{}

func (c *Coffee) Drink() {
	fmt.Println("Drinking coffee")
}

type Tea struct{}

func (t *Tea) Drink() {
	fmt.Println("Drinking tea")
}

type Cake struct{}

func (c *Cake) Eat() {
	fmt.Println("Eating cake")
}

type Cookie struct{}

func (c *Cookie) Eat() {
	fmt.Println("Eating cookie")
}

func CreateVoucher(food Food, drink Drink) Voucher {
	return Voucher{Food: food, Drink: drink}
}


func main() {
	// Create voucher with coffee and Cookie
	// Coffee and Cookie are not related

	// v := Crea{food: &Cookie{}, drink: &Coffee{}}
	v := CreateVoucher(&Cookie{}, &Coffee{})
	fmt.Print(v)
}
