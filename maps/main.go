package main

import "fmt"

func main()  {
    a := make(map[int]string)
    b := new(map[int]string)
    a[1] = "a"
    a[2] = "b"
    v,f := a[2]
    z := map[string]int{
      "a" : 1,
      "b" : 2,
      "c" : 3,
    }
    for key, value := range z{
      fmt.Println(key, " -> ", value)
    }
    fmt.Println( a, "\n", b, v, f ,"\n", z)
}
