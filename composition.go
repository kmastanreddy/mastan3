package main
import "fmt"
type Person struct {
 id int
 name string
}
type Employee struct {
 Person
 desg string
 salary float64
}
 func main() {
     e:= Employee{}
     e.id=101
     e.name="Srikar"
     e.desg="Manager"
     e.salary=99999.99
    fmt.Printf("%v",e)
}  




