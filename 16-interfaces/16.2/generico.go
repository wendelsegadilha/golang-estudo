package main

import "fmt"

func gnerica(interf interface{})  {
	fmt.Println(interf)
}

func main()  {
	gnerica("Wendel")
	gnerica(10)
	gnerica(true)
	gnerica(3.14)
}