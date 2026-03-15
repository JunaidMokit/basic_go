package main

import "fmt"

func main() {
   

mark:=20

if mark>50{
	fmt.Println("You are passed with good mark")
} else if mark>33{
	fmt.Println("You are passed ")
} else{
	fmt.Println("You are not eligible")
}

}