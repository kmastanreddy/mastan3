package main
import  (
   "fmt"
   "os"
)
func main() {
  argCount:=len(os.Args[1:])
  fmt.Printf("Total number of args :%d\n",argCount)
 for i,a:=range os.Args[1:] {
    fmt.Printf("Index %d value %s\n",i+1,a)

  } 
}

  




