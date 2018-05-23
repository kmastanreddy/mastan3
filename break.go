package main
import "fmt"
func main() {
  a:=1
   for a<10 {
     	fmt.Println(a)
        a++
        if(a>5){
          break
        }
  }
}