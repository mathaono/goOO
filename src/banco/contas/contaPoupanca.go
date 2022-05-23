package contas

import "banco/clientes"

type ContaPoupanca struct {
	Nome          clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Saque(valor float64) string {
	if valor <= c.saldo && valor > 0 {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "Operação indisponível"
	}
}

func (c *ContaPoupanca) Deposita(valor float64) string {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Operação indisponível"
	}
}

func (c *ContaPoupanca) MostraSaldo() (string, float64) {
	return "Saldo Atual: ", c.saldo
}
