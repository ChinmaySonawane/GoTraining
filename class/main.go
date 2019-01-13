package main

import ("fmt"
        "time"
)

type student struct {

    name string
    age int
    entry time.Time
}

func (this student)add(list []student)[]student{

    var x student
    fmt.Scan(&x.name)
    fmt.Scan(&x.age)
    x.entry = time.Now()
    fmt.Println(x)
    return append(list, x)
}

func (this student)show(list []student){

    for _,v := range list {
      fmt.Println(v.name, "\t \t", v.age, "\t ", v.entry.Format("2006-01-02 15:04:05"))
    }

}

func main()  {
  var x student
  list := make([]student,0)
  for i := 0;i < 3;i++ {
    list = x.add(list)
  }
  x.show(list)
}
