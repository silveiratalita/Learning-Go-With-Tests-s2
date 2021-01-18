package main

import "fmt"

func Ola()string{
	return "Olá mundão"
}
func OlaYou(name string) string{
	return "Olá, "+name
}

func main(){
	fmt.Println(Ola())
	fmt.Println(OlaYou("Talita"))
}