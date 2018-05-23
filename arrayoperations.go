package main
import "fmt"
func main() {
  a := []int{1,2,3}
  b := []int{4,5,6}
  c := append(a,b...)
  fmt.Println(c)

  i := []interface{}{10,20,30}
  j := []interface{}{"Java","Oracle","Golang"}
  k := append(i,j...)
  fmt.Println(k)

  x:=[...]int{7,8,9}
  y:=[...]int{10,11,12}
  var z [len(x)+len(y)] int
  copy(z[:],x[:])
  copy(z[len(x):],y[:])
  fmt.Println(z)
 
}

  




