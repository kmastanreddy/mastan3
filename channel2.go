package main
import   (
   "fmt"
   "time"
   "sync"
)
var wait sync.WaitGroup
   func ping(done chan string) {
   fmt.Println("routine with channel writing")
   done <-"Chanel in routines"  // writing into channel
   wait.Done()
}
func print(done chan string) {
     fmt.Println("routine with channel reading")
     msg:= <- done  // writing into channel
     fmt.Println(msg)
     time.Sleep(time.Second*1)
     wait.Done();
}
func main() {
   var done chan string =make(chan string)
   wait.Add(2)
   go ping(done)
   go print(done)
   wait.Wait()
 }

