package main

import (
	"fmt"
)

type person struct {
	id   int
	name string
	age  int
}

func main() {
	conn := InitDB()
	defer conn.Close()
	// Query 返回的是Rows
	rows, err := conn.Query("select * from test")
	if err != nil {
		fmt.Println(err.Error())
	}
	// 将result 解析成 结构体对象
	var persons = make([]person, 0)
	for rows.Next() {
		var a person
		rows.Scan(&a.id, &a.name, &a.age)
		persons = append(persons, a)
	}
	fmt.Println(persons)
}
