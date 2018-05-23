package main
import "fmt"
func main() {
  cities := []string{"chennai", "Hyderabad"}
  othercities := []string {"Banglore","Pune"}
  cities = append(cities,othercities...)
  fmt.Printf("%q",cities)
 
}

  




