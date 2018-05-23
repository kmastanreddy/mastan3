package main
import "fmt"
func main() {
  var a int = 0
   for a<10 {
     if a == 5 {
        // a=a+1;
          continue;
        }
     a++
     fmt.Println(a)
  }
}