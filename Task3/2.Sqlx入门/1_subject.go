package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func initDB() (db *sqlx.DB, e error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/youbase?charset=utf8mb4&parseTime=True"
	db, e = sqlx.Connect("mysql", dsn)
	return db, e
}

type Employees struct {
	Id         int     //`db:"id"`
	Name       string  //`db:"name"`
	Department string  //`db:"department"`
	Salary     float32 //`db:"salary"`
}

func seekEmployees(db *sqlx.DB) {
	emps := make([]Employees, 0, 10)
	r, _ := db.Queryx("SELECT * FROM employees WHERE department = \"技术部\";")
	for r.Next() {
		emp := Employees{}
		r.Scan(&emp.Id, &emp.Name, &emp.Department, &emp.Salary)
		emps = append(emps, emp)
	}
	fmt.Println(emps)
}

func seekMostSalary(db *sqlx.DB) {
	emp := Employees{}
	r, _ := db.Query("SELECT * FROM employees ORDER BY salary DESC")
	r.Next()
	r.Scan(&emp.Id, &emp.Name, &emp.Department, &emp.Salary)
	fmt.Println(emp)
}

func main() {
	db, e := initDB()
	if e != nil {
		panic(e)
	}
	seekEmployees(db)
	seekMostSalary(db)

	test2Main(db)
}
