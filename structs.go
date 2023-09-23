package main
import "fmt"

type user struct {
	name string
	phone int
}

// nested struct
type msg struct {
	message string
	sender user
	recipient user
}

type car struct {
	make string
	model string
	// anonymous struct
	Wheel struct {
		radius int 
		material string
	}
}

// embedded struct
type ownedCar struct {
	user
	make string
	model string
}

type cuboid struct {
	length int 
	width int 
	height int
}

// take struct as self
func (c cuboid) volume() int {
	return c.length * c.width * c.height
}

func main() {
	m := user {
		phone: 148255510961,
		name: "Sumit",
	}
	fmt.Println(m)

	m2 := msg {}
	m2.message = "Thanks for signing up"
	m2.sender.phone = 1234567890
	m2.sender.name = "Boot dev"
	m2.recipient.phone = 148255510961
	m2.recipient.name = "Sumit"
	fmt.Println(m2)

	// anonymous structs
	myCar := struct {
		make string
		model string
	} {
		make: "tesla",
		model: "model 3",
	}
	fmt.Println(myCar)

	// embedded structs
	myCar2 := ownedCar {
		make: "bugatti",
		model: "chiron",
		user: user {
			name: "Sumit",
			phone: 148255510961,
		},
	}
	fmt.Println(myCar2)

	// struct methods
	c := cuboid {
		length: 7,
		width: 5,
		height: 16,
	}
	fmt.Println(c.volume())
}
