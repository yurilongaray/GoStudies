package main

import (
	"fmt"
	"math"
)

type person struct {
	name   string
	age    int
	salary float64
}

type dentist struct {
	person
	totalConsultation int
}

type driver struct {
	person
	kmDriven int
}

type humam interface {
	hello(message string)
}

func (x dentist) hello(message string) {

	fmt.Println("My name is ", x.name, "I made", x.totalConsultation, message)
}

func (x driver) hello(message string) {

	fmt.Println("My name is ", x.name, "I drove ", x.kmDriven, message)
}

func humamBeingStuff(h humam) {

	switch h.(type) {
	case dentist:
		h.hello("How u doing?")
	case driver:
		h.hello("Where u going?")
	}
}

type square struct {
	size float64
}

type circle struct {
	raio float64
}

type figure interface {
	calculateArea() float64
}

func (c circle) calculateArea() float64 {

	return ((2 * math.Pi) * c.raio)
}

func (s square) calculateArea() float64 {

	return s.size * s.size
}

func infoArea(f figure) {

	fmt.Println(f.calculateArea())
}

//example using pointer into receiver
func (d *driver) walk() {

	fmt.Println("Using pointer into receiver", d.kmDriven)
}

type humamActions interface {
	walk()
}

func main() {

	dentist1 := dentist{
		person: person{
			name: "Loki",
		},
		totalConsultation: 99,
	}

	driver1 := driver{
		person: person{
			name: "Thor",
		},
		kmDriven: 222000,
	}

	humamBeingStuff(dentist1)
	humamBeingStuff(driver1)

	circle1 := circle{raio: 2}
	square1 := square{size: 2}

	infoArea(circle1)
	infoArea(square1)

	humamActions.walk(&driver1) //example using pointer

	// this is a shortcut for (&driver1).walk()
	driver1.walk()

	// this is the correct way to call props from struct (but is unusual)
	(&driver1).walk()

	// interface as generic type (This is the most generic type). You can use switch .type() to get the interface values type
	var stuff interface{}
	var stuff2 interface{}

	stuff = 1
	stuff2 = "LOL"

	fmt.Println(stuff, stuff2)
}
