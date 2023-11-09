package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kletka bool

var totalX = 50
var totalY = 20
var lifePercent = 50

func (table room) GetLifeArround(xstart int, ystart int) int {
	totalLife := 0
	for i := xstart - 1; i <= xstart+1; i++ {
		if i < 0 {
			continue
		}
		if i >= totalX {
			continue
		}

		for y := ystart - 1; y <= ystart+1; y++ {
			if y < 0 {
				continue
			}
			if y >= totalY {
				continue
			}
			if (xstart == i) && (ystart == y) {
				continue
			}
			if table[y][i] {
				totalLife++
			}
		}
	}
	return totalLife
}

func (table *room) makeLife() {

	for ind1, line := range *table {
		for ind2, el := range line {
			life := false
			sm := rand.Intn(2)
			if sm == 1 {
				life = true
			}
			_ = el
			line[ind2] = life

		}
		fmt.Println()
		_ = ind1
	}

}

func (r *room) Inizialize() {
	var line []bool
	for i := 0; i < totalY; i++ {
		line = make([]bool, totalX)
		*r = append(*r, line)
	}
}

type room [][]bool

func (someRoom room) showAll() {
	for _, el := range someRoom {
		for _, el2 := range el {
			if el2 {
				fmt.Print("0")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
	time.Sleep(time.Millisecond * 100)
}

func (someRoom room) CheckGeneration() *room {
	var newGen room
	newGen.Inizialize()
	for y, el := range someRoom {
		for x, _ := range el {
			life := someRoom.GetLifeArround(x, y)
			newlife := someRoom[y][x]
			if life == 3 { // есть
				newlife = true
				//continue
			}
			if life == 2 { // не меняется
				//continue
			}
			if (life < 2) || (life > 3) { //умирает
				newlife = false
				//continue
			}
			newGen[y][x] = newlife
			_ = newlife
		}
	}
	return &newGen
}

func main() {
	// сначало 1е поколение - случайное заполнение
	// В пустой если рядом есть 3 живые - идет жизнь
	// если меньше 2х или больше 3х клетка умирает
	// если ничего не меняется в течении n ходов все прекращается
	var r room
	r.Inizialize()
	r.makeLife()
	r.showAll()
	for i := 0; i < 1000; i++ {
		r = *r.CheckGeneration()
		r.showAll()
	}
	r.showAll()

	//fmt.Printf("r: %v\n", r[10][49])

}
