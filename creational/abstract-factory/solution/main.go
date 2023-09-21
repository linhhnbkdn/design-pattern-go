package main

import (
	"errors"
	"fmt"
)

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

type VoucherAbstractFactory interface {
	GetDrink() Drink
	GetFood() Food
}

type CoffeeMorningVoucherFactory struct{}
func (CoffeeMorningVoucherFactory) GetDrink() Drink { return &Coffee{} }
func (CoffeeMorningVoucherFactory) GetFood() Food { return &Cake{} }


type DrinkEveningVoucherFactory struct{}
func (DrinkEveningVoucherFactory) GetDrink() Drink { return &Tea{} }
func (DrinkEveningVoucherFactory) GetFood() Food { return &Cookie{} }

func GetVoucherFactory(campaignName string) (VoucherAbstractFactory, error){
	if campaignName == "CoffeeMorning" {
		return CoffeeMorningVoucherFactory{}, nil
	} else if campaignName == "DrinkEvening" {
		return DrinkEveningVoucherFactory{}, nil
	}
	return nil, errors.New("Wrong campaign name passed")
}

func GetVoucher(factory VoucherAbstractFactory) Voucher {
	return Voucher{
		Drink: factory.GetDrink(),
		Food: factory.GetFood(),
	}
}
func main() {
	voucherFactory, err := GetVoucherFactory("CoffeeMorning")
	if err != nil {
		fmt.Println(err)
		return
	}
	voucher := GetVoucher(voucherFactory)
	voucher.Drink.Drink()
	voucher.Food.Eat()

}
