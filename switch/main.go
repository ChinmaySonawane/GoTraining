package main

import "fmt"

func main()  {
    var x ,v = 10, 12
    switch  {
    case x == 10  ,v == 12  : fmt.Println("case 2")
    default       : fmt.Println("default")
    }
}
