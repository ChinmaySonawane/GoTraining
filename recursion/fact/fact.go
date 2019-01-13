package fact

func Fact(no int)int  {
    if no==1 {
      return 1
    }
    return Fact(no - 1) * no
}
