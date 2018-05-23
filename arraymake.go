package main
import  "fmt"

func main() {
 a:=make([]int,5)
  fmt.Println(a,len(a),cap(a))
 b:=make([]int,5)
    fmt.Println(b,len(b),cap(b))
  c:=b[:2]
    fmt.Println(c,len(c),cap(c))
  d:=b[2:5]
       fmt.Println(d,len(d),cap(d))

 }

 