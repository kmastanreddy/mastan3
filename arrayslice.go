package main
import "fmt"
func main() {
  courses:=[5] string {
   "java",
    ".net",
    "golang",
    "scala",
    "reactjs",
   }
   fmt.Println(courses)
   slice1:=courses[0:2]
   slice2:=courses[3:5]
   fmt.Println(" slices.....")
   fmt.Println(slice1,slice2)
   fmt.Println(courses)


  }
