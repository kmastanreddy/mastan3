package main
import   (
   "fmt"
 )
func p1() {
        fmt.Println("Hi")
}
func p2() {
        fmt.Println("Hello")
}
func main() {
   p2()
   p2()
  defer p1()
  
 }

