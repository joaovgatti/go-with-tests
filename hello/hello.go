package main

import "fmt"


const englishPrefix = "Hello, "
const spanishPrefix = "Hola, "
const portuguesePrefix = "Olá, "

func HelloUser(name, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}


func greetingPrefix(language string) (prefix string) {
	switch language {
		case "spanish":
			prefix = spanishPrefix
		case "portuguese":
			prefix = portuguesePrefix
		default:
			prefix = englishPrefix

	}
	return

}

func main(){
	fmt.Println(HelloUser("Joao","english"))
}