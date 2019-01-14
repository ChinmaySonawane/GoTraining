package main

import "fmt"

func main() {

    s := make([]int, 5)

    for i := 0; i < cap(s); i++ {
        s[i] = i
    }

    fmt.Println(s)

    var p = make([]int, 10)
    copy(p,s)//p[10] and s[5]
    p=append(p,s...)
    fmt.Println(p[:])

    twoD := make([][]int, 3)
    for i := 0; i < cap(twoD); i++ {

      twoD[i] =make([]int, i+1)
      for j := 0; j < cap(twoD[i]); j++ {
        twoD[i][j] = i+j
      }
    }
    twoD[1] = nil
    fmt.Println(twoD)
}
