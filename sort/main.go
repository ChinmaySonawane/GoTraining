package main

import (
	"fmt"
	"sort"
)

type MyInt []string

func (m MyInt) Len() int {
	return len(m)
}

func (m MyInt) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m MyInt) Less(i, j int) {
	return m[i] > m[j]
}

func main() {
	var a = MyInt{"chinmay", "anuj", "pradip", "p", "a"}
	fmt.Printf("%v", a)
	sort.Sort(MyInt(a))

}
