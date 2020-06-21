package main

import "fmt"

func greeting(name string)string  {
	return "hellow "+ name
}

func main()  {
	fmt.Print(greeting("dan"))
}