package main
import  "fmt"
func main() {
   var a chan int
   if a==nil {
  fmt.Println("channel is nil")
   a=make(chan int)
    fmt.Printf(" type of a is %T",a)
}
}
