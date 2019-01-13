package main

import "fmt"

func main()  {
  p :=[5]string{"a", "b", "c","",""}
  //var q [5]string
  var x [5][5]int
  var v int
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            x[i][j] = v
            v++
        }
    }
    var sum int
  //  copy(q, p) only
    for i,_ := range x {
      for _,val := range x[i] {
        sum += val
        fmt.Println(i, val)
      }
    }

    fmt.Println(x, "\n ", len(x), cap(x) ,"\n", p,"\n sum ",sum)//, q)
}
