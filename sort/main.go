package main

import ("fmt"
		"sort"
)

type MyInt []int

func (m MyInt) Len() int {
	return len(m)
}

func (m MyInt) Swap(i, j int)  {
	m[i], m[j] = m[j], m[i] 
}

func (m MyInt) Less(i, j int) bool {
	return m[i] > m[j]
}

func main()  {
	var a =MyInt{1,2,3,4,5,6,7}
	fmt.Printf("%v",a)
	sort.Sort(MyInt(a))
	fmt.Printf("%v",a)
}