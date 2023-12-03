package main

import "fmt"

type Item interface{
	Cost() float32
}

type RealBox struct {
	Name string
	Price float32
}

func (realBox RealBox) Cost() float32 {
	return realBox.Price
}

type RootBox struct {
	Name string
	children []Item
}

func (rootBox RootBox) Cost() float32 {
	var cost float32 = 0
	for _, child := range rootBox.children {
		cost += child.Cost()
	}
	return cost
}

func CreateItem() Item {
	return RootBox{
		Name: "Root",
		children: []Item{
			RealBox{
				Name: "Child 1",
				Price: 1,
			},
			RealBox{
				Name: "Child 2",
				Price: 2,
			},
			RootBox{
				Name: "Root A",
				children: []Item{
					RealBox{
						Name: "Child A.1",
						Price: 1.1,
					},
				},
			},
		},
	}
}

func main() {
	root := CreateItem()
	fmt.Printf("%f\n", root.Cost())
}

