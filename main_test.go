package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestAddPozitive(t *testing.T) {
	a, err := Add(1, 2)
	if err != nil {
		t.Error("Unexpected error")
	}

	if a != 3 {
		t.Errorf("Sum expected to be 3; got %d", a)
	}
}

func TestAddNegative(t *testing.T) {
	fmt.Println("asdf")
	_, err := Add(-1, -2)
	if err == nil {
		t.Error(" first arg negative - expected error not be nil	")
	}
	_, err = Add(1, -2)
	if err == nil {
		t.Error(" second arg arg negative - expected error not be nil	")
	}
	_, err = Add(-1, -2)
	if err == nil {
		t.Error(" all arg arg negative - expected error not be nil	")
	}

}

func TestAllZero(t *testing.T) {
	fmt.Println("All 00")
	_, err := Add(0, 1)
	if err != nil {
		errors.New("Lef args is 0. Return Must be nil ")
	}
	_, err = Add(1, 0)
	if err != nil {
		errors.New("Right args is 0. Return Must be nil ")
	}
	_, err = Add(0, 0)
	if err != nil {
		errors.New("Both args is 0. Return Must be nil ")
	}

}
