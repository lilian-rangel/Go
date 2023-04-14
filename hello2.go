package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Lilian"
	var dias = 3
	var versao = 3.1
	fmt.Println("Olá, sra", nome, "Este é seu", dias, "dia aprendendo go")
	fmt.Println("Este é a versão", versao, "do seu código")
	fmt.Println("O Go consegue detectar o tipo de variável então não precisa ser declarada explicitamente. Isso é chamado de inferência de variável")

	fmt.Println("O tipo da variável nome é", reflect.TypeOf(nome))
}
