package main

import "fmt"

type Student struct{
	name string
	age int
	schoolName string
}

func (s Student) String() string {
	
	return fmt.Sprintf("name : %v \nage : %v\nschool name : %v",s.name,s.age,s.schoolName)
}

func main()  {
	
	v := Student{"chinmay", 21, "ramanbaug"}
	fmt.Printf("%v",v)
}