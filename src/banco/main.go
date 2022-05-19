package main

import (
	"banco/contas"
	"fmt"
)

func main() {

	contaMatheus := contas.ContaCorrente{nome: "Matheus", saldo: 250}
	contaJose := contas.ContaCorrente{nome: "Jose", saldo: 750}
	fmt.Println(contaMatheus)
	fmt.Println(contaJose)

	contaJose.Transfere(250, &contaMatheus)

	fmt.Println(contaMatheus)
	fmt.Println(contaJose)
}
