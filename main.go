package main

import "fmt"

func main() {
	fmt.Println(SayHello())
}

func SayHello() string {
	return fmt.Sprint("Hello")
}
