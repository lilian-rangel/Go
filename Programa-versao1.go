package main

import (
	"fmt"
)

func main() {
	nome := "Lilian"

	fmt.Println("Olá, sra", nome)
	fmt.Println("Escolha a opção que deseja executar")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("A opção escolhida foi:", comando)

}
