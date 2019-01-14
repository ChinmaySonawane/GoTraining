package main

import "fmt"

type Student struct{
	name string
	age int
}

func NewStudent(name string,age int) *Student {
	
	s := new(Student)
	s.name = name
	s.age = age
	return s
}

func main()  {
	v := Student{"chinmay",21}
	fmt.Println(v)
}