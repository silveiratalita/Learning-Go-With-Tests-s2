package main

import "fmt"

const prefixHello = "Olá, "

func Ola()string{
	return prefixHello + "mundão"
}
func OlaYou(name string) string{
	return prefixHello + name
}

func main(){
	fmt.Println(Ola())
	fmt.Println(OlaYou("Talita"))
}