package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Saque(valorBoleto)
}

type verificarConta interface {
	Saque(valor float64) string
}

func main() {

	clienteMatheus := clientes.Titular{"Matheus", "123.123.123.12", "Dev"}
	contaMatheus := contas.ContaPoupanca{Nome: clienteMatheus, NumeroAgencia: 123, NumeroConta: 777}
	contaMatheus.Deposita(1000)
	fmt.Println(contaMatheus.MostraSaldo())

	PagarBoleto(&contaMatheus, 200)
	fmt.Println(contaMatheus.MostraSaldo())
}
