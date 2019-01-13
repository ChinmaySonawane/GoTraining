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

  //  copy(q, p) only

    fmt.Println(x, "\n ", len(x), cap(x) ,"\n", p)//, q)
}
