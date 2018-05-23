package main
import  (
       "fmt"
       "time"
       "sync"
)
var wg =sync.WaitGroup{}
func main() {
   wg.Add(2)
   go f1()
   go f2()
   wg.Wait()
}
func f1(){
 for i:=0;i<=10;i++ {
    fmt.Println("function1",i)
    time.Sleep(time.Duration(5*time.Millisecond))
  }
   wg.Done()
}
func f2(){
 for i:=0;i<=10;i++ {
    fmt.Println("function2",i)
    time.Sleep(time.Duration(10*time.Millisecond))
  }
   wg.Done()

}
