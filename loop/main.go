package main

import "fmt"

func main() {
  //fibonacci()
  /*a,b := 1,2
  a,b = b,a
  fmt.Println("a ", a, "b ", b)
*/
    for i := 10; i > 0 ; i-- {
        fmt.Println("i ", i)
        if i % 2 == 0 {
          continue
        }
        i--
    }

    for i := 10; i > 0 ; i-- {
        fmt.Println("i ", i)
        if i == 8 {
          break//return
        }
    }
    fmt.Println("bye")
}
