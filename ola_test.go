package main

import "testing"

func TestOla(t *testing.T){
	resultado:=Ola()
	esperado:="Olá mundão"
	if resultado != esperado {
	t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}