package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	ford := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	sClass := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println("=================== Truck")
	fmt.Println(ford)
	fmt.Println(ford.doors, ford.color, ford.fourWheel)
	fmt.Println("=================== Sedan")
	fmt.Println(sClass)
	fmt.Println(sClass.doors, sClass.color, sClass.luxury)
}
