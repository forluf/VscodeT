package mypack

// package main

import (
	"fmt"
	"reflect"
)

type Spell interface {
	// название заклинания
	Name() string
	// характеристика, на которую воздействует
	Char() string
	// количественное значение
	Value() int
}

// CastReceiver — если объект удовлетворяет этом интерфейсу, то заклинание применяется через него
type CastReceiver interface {
	ReceiveSpell(s Spell)
}

func CastToAll(spell Spell, objects []interface{}) {
	for _, obj := range objects {
		CastTo(spell, obj)
	}
}

func CastTo(spell Spell, object interface{}) {

	if recv, ok := object.(CastReceiver); ok {
		recv.ReceiveSpell(spell)
		fmt.Printf("recv: %v\n", recv)
		fmt.Printf("recv: %v\n")
		return
	}

	//fmt.Println(spell, " ", object)
	elem := reflect.ValueOf(object).Elem()
	size := reflect.ValueOf(object).Elem().NumField()
	for i := 0; i < size; i++ {
		field := elem.FieldByName("Health")
		if !field.IsValid() {
			//fmt.Printf("field invalid: %v\n", field)
			break
		}
		if field.CanSet() {
			//fmt.Println("Can set")
		}
		if !field.IsZero() {
			//fmt.Printf("field not zero: %v\n", field.Kind())
			field.SetInt(10000)
		}
		//fmt.Println(object)
	}

}

type spell struct {
	name string
	char string
	val  int
}

func newSpell(name string, char string, val int) Spell {
	return &spell{name: name, char: char, val: val}
}

func (s spell) Name() string {
	return s.name
}

func (s spell) Char() string {
	return s.char
}

func (s spell) Value() int {
	return s.val
}

type Player struct {
	name   string
	health int
}

func (p *Player) ReceiveSpell(s Spell) {
	if s.Char() == "Health" {
		p.health += s.Value()
	}
}

type Zombie struct {
	Health int
}

type Daemon struct {
	Health int
}

type Orc struct {
	Health int
}

type Wall struct {
	Durability int
}

func main() {

	player := &Player{
		name:   "Player_1",
		health: 100,
	}

	enemies := []interface{}{
		&Zombie{Health: 1000},
		&Wall{Durability: 100},
	}

	CastToAll(newSpell("fire", "Health", -50), append(enemies, player))
	//	CastToAll(newSpell("heal", "Health", 190), append(enemies, player))
	//fmt.Println(player)
	for _, v := range enemies {
		fmt.Printf("v: %v\n", v)
	}

}
