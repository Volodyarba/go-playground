package main

import "fmt"

func main() {
    var motivation int

    fmt.Print("lvl your motivation 1-10: ")
    fmt.Scanln(&motivation)

   if motivation < 4 {
	fmt.Println("low motivation")
   } else if motivation < 7 {
	fmt.Println("mid motivation")
   } else if motivation < 11 {
	fmt.Println("hard motivation")
   } else {
	fmt.Println("no currect position")
   }
}