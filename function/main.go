package main

import "fmt"

func Sum(nums ...int)(sum int) {

  for _,v := range nums{
    sum +=v
    }
  return
}

func main()  {

  fmt.Println("sum ",Sum(1,2,3,4))
}
