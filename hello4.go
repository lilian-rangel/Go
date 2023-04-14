package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Lilian"
	dias := 3
	var versao = 3.4
	fmt.Println("Olá, sra", nome, "Este é seu", dias, "dia aprendendo go")
	fmt.Println("Este é a versão", versao, "do seu código")
	fmt.Println("O Go consegue detectar o tipo de variável então não precisa ser declarada explicitamente.\nIsso é chamado de inferência de variável.\nTem também o atribuidor de variáveis curto onde não é preciso escrever 'var' antes mas utilizar essa combinação :=")

	fmt.Println("É possível descobrir o tipo de variável importando 'reflect'.\nO tipo da variável nome é", reflect.TypeOf(nome))
}
