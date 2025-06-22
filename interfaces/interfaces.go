package interfaces

import "fmt"

type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Меня зовут " + p.Name
}

func Interface() {
	var s Speaker = Person{Name: "Шляпа"}
	fmt.Println(s.Speak())
	defer fmt.Println("end of interface func")
}
