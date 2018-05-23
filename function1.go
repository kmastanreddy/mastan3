package main
import "fmt"
  func fun() int { //return type of function
  return 123    //function body
}
func main() {
  x:=fun()  //function call := type inference

  fmt.Println(x) //return value from function
  fmt.Println(fun())  /inline function call
}