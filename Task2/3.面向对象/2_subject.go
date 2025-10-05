package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID int
}

func (self *Employee) PrintInfo() {
	fmt.Println("姓名：", self.Name)
	fmt.Println("年龄：", self.Age)
	fmt.Println("员工编号：", self.EmployeeID)
	fmt.Println("————————————————")
}

func main() {
	employees := []Employee{
		Employee{
			Person:     Person{Name: "小明", Age: 19},
			EmployeeID: 100011,
		},
		Employee{
			Person:     Person{Name: "小王", Age: 20},
			EmployeeID: 100012,
		},
		Employee{
			Person:     Person{Name: "小张", Age: 18},
			EmployeeID: 100013,
		},
	}
	for _, v := range employees {
		v.PrintInfo()
	}
}
