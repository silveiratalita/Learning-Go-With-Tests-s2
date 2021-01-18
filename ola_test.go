package main

import "testing"

func TestOla(t *testing.T){
	result:=Ola()
	expect:="Olá mundão"
	if result != expect {
	t.Errorf("result '%s', expect '%s'", result, expect)
	}
}

func TestOlaYou(t *testing.T){
	result:= OlaYou("Talita")
	expect:= "Olá, Talita"
	if result != expect{
		t.Errorf("result '%s',expect '%s'",result, expect)
	}
}