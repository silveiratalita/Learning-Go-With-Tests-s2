package main

import "fmt"



func Ola()string{
	return "Olá, mundão"
}
func OlaYou(name string, language string) string{
	prefixHello := func(language string )string {
		switch language {
		case "pt":
			return "Olá, "
		case "es":
			 return"Hola, "
		default:
			return "Hello, "
		}
	}
	if name == ""{
		return prefixHello(language) + "mundão"
	}
	return prefixHello(language) + name
}

func main(){
	fmt.Println(Ola())
	fmt.Println(OlaYou("Talita","es"))
}