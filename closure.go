package main
import "fmt"
func main() {
    sayhello:= func() string {
    return "Hello"
    }  
    sum := func(a int,b int) int {
     return a*a+b*b;
}
    fmt.Println(sayhello())
    fmt.Println(sum(10,20))
    fmt.Println("welcome to closure")
}


