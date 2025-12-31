package main

import "fmt"

func main() {
   // fmt.Println("Hello, World!")
   var x int =10
   y:=20
   sum:=x+y
   fmt.Println("Result is=",sum)

   if sum>20 {
    fmt.Println("Perfact Sum")
   } else if sum <20 && sum >15 {
    fmt.Println("Moderate Sum")
   } else {
    fmt.Println("Not Good")
   }

   for i:=0;i<5;i++ {
    fmt.Println(i)
   }
}