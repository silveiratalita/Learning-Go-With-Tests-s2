package main

import "testing"


func TestOlaYou(t *testing.T){

	verifyResultAndExpect:=func(t *testing.T, result, expect string){
		t.Helper()
		if result != expect{
			t.Errorf("result '%s',expect '%s'",result, expect)
		}
	}

	t.Run("Diz Olá para as pessoas em portugês", func(t *testing.T){
		result:= OlaYou("Talita","pt")
		expect:= "Olá, Talita"
		verifyResultAndExpect(t, result,expect)
	})
	t.Run("Diz Olá para as pessoas em espanhol", func(t *testing.T){
		result:= OlaYou("Talita","es")
		expect:= "Hola, Talita"
		verifyResultAndExpect(t, result,expect)
	})
	t.Run("Diz Olá mundão quando a string vem vazia", func(t *testing.T){
		result:= OlaYou("","pt")
		expect:= "Olá, mundão"
		verifyResultAndExpect(t, result,expect)
	})
}

func TestOla(t *testing.T){
	result:= Ola()
	expect:= "Olá, mundão"
	if result != expect{
		t.Errorf("result '%s',expect '%s'",result, expect)
	}
}