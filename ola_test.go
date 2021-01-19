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
func TestAdd(t *testing.T){
	verifySumAndExpect:=func(t *testing.T, sum, expect int){
		t.Helper()
		if sum != expect{
			t.Errorf("result '%d',expect '%d'",sum, expect)
		}
	}
	t.Run("Adiciona 2 + 2", func(t *testing.T){
		sum:= AddValue(2, 2)
		expect:= 4
		verifySumAndExpect(t, sum,expect)
	})
}