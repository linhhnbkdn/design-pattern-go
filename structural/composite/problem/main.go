package main

type Item struct {
	Name string
	Price float64
	children []Item
}

func (item Item) Cost () float64 {
	total := item.Price
	for _, child := range item.children {
		total += child.Cost()
	}
	return total
}

func CreateItem () Item {
	return Item{
		Name: "Root",
		Price: 0,
		children: []Item{
			{
				Name: "Child 1",
				Price: 1,
				children: []Item{
					{
						Name: "Child 1.1",
						Price: 1.1,
						children: []Item{},
					},
				},
			},
			{
				Name: "Child 2",
				Price: 2,
				children: []Item{},
			},
		},
	}
}

func main() {
	root := CreateItem()
	println(root.Cost())
}


// structure này không tốt vì cái root có thể có nhiều child, child có thể có nhiều child
// Hop root thường không có giá trị, nó chỉ là một cái container để chứa các child
