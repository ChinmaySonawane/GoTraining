package main

import "fmt"

func closure(v int) func() int {

  var no = v

  return func() int {
    no++
    return no
  }

}

  func main()  {
      f := closure(10)
      fmt.Println(f(),f(),f())
      m := closure(5)
      fmt.Println(m(),m(),f())

}
