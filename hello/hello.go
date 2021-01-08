package main

import "fmt"

func HelloWorld() string {
	return "Hello World"
}

func HelloUser(name string) string {
	return "Hello " + name
}

func main(){
	fmt.Println(HelloWorld())
	fmt.Println(HelloUser("Joao"))
}