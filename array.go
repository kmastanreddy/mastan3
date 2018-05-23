package main
import "fmt"
func main() {
  var x[5] int 
  var i,j int
  for i=0;i<5;i++ {
    fmt.Scanln(&x[i]) //read array values at runtime
     //x[i]=i+10 //10,11,12,13,14
  }
  for j=0;j<5;j++ {
     fmt.Printf("Index[%d] = %d\n",j,x[j])
  }
  s1:=x[1:3]    //slicing
  fmt.Println(s1)
}