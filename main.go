package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Cow struct {
	name string
}

func (c Cow) move() string {
	switch rand.Intn(4) {
	case 0:
		return "лево"
	case 1:
		return "право"
	case 2:
		return "назад"
	case 3:
		return "вперед"
	default:
		return "стоим на месте"
	}
}

func (c Cow) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "сено"
	case 1:
		return "траву"
	case 2:
		return "корм"
	case 3:
		return "не хочет есть"
	default:
		return "ищем еду"
	}
}

type Bird struct {
	name string
}

func (c Bird) move() string {
	switch rand.Intn(4) {
	case 0:
		return "вверх"
	case 1:
		return "садимся"
	case 2:
		return "взлетаем"
	case 3:
		return "кругом"
	default:
		return "парим на месте"
	}
}

func (c Bird) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "солому"
	case 1:
		return "воду"
	case 2:
		return "зерно"
	case 3:
		return "не хочет есть"
	default:
		return "ищем еду"
	}
}

func (c Bird) String() string {
	return ("птичка " + c.name)
}
func (c Cow) String() string {
	return ("корова " + c.name)
}

type Animal interface {
	move() string
	eat() string
}

func makeSomething(a Animal) {
	action := rand.Intn(2)
	switch action {
	case 0:
		fmt.Printf("%v двигеается - %v \n", a, a.move())
	case 1:
		fmt.Printf("%v кушает - %v \n", a, a.eat())
	default:
		fmt.Println("Ничего не делают")
	}
}

func main() {
	Sunrise := 8
	Sunset := 20
	maxday := 3

	var somecow, popugai Animal

	somecow = Cow{name: "Zorka"}
	popugai = Bird{name: "Popug"}

	var someAnimal []Animal

	someAnimal = append(someAnimal, somecow, popugai)

	for day := 0; day < maxday; day++ {
		fmt.Println("Day - " + strconv.Itoa(day))
		for hour := 0; hour < 24; hour++ {
			if hour < Sunrise || hour > Sunset {
				fmt.Println("Все спят")
			} else {
				makeSomething(someAnimal[rand.Intn(len(someAnimal))])
			}
			time.Sleep(time.Second / 3)
		}
	}
}
