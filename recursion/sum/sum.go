package sum

func Sum(no int)int  {
    if no == 1 {
        return no
    }
    return Sum(no - 1) + no
}
