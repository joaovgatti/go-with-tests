package main

import "fmt"


const englishPrefix = "Hello, "

func HelloUser(name string) string {
	if name == ""{
		name = "World"
	}
	return englishPrefix + name
}

func main(){
	fmt.Println(HelloUser("Joao"))
}