package main

import "fmt"

type Shape interface {
	Area()
	Perimeter()
}
type Base struct {
	name string
}

type Rectangle struct {
	Base
}

func (self *Rectangle) Area() {
	fmt.Println(self.name + ":Area")
}
func (self *Rectangle) Perimeter() {
	fmt.Println(self.name + ":Perimeter")
}

type Circle struct {
	Base
}

func (self *Circle) Area() {
	fmt.Println(self.name + ":Area")
}
func (self *Circle) Perimeter() {
	fmt.Println(self.name + ":Perimeter")
}

func main() {
	shapes := []Shape{&Rectangle{Base: Base{name: "矩形"}}, &Circle{Base: Base{name: "圆形"}}}
	for _, v := range shapes {
		v.Area()
		v.Perimeter()
	}
}
