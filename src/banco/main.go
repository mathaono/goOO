package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaMatheus := contas.ContaCorrente{Nome: "Matheus", Saldo: 250}
	contaJose := contas.ContaCorrente{Nome: "Jose", Saldo: 750}
	fmt.Println(contaMatheus)
	fmt.Println(contaJose)

	contaJose.Transfere(250, &contaMatheus)

	fmt.Println(contaMatheus)
	fmt.Println(contaJose)
}
