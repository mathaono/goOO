package contas

import "banco/clientes"

type ContaCorrente struct {
	Nome          clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Saque(valor float64) string {
	if valor <= c.saldo {
		if valor > 0 {
			c.saldo -= valor
			return "Saque realizado com sucesso"
		} else {
			return "Não é permitido valor negativo"
		}
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Deposita(valor float64) string {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso"
	} else {
		return "Impossível depositar valor negativo"
	}
}

func (c *ContaCorrente) Transfere(valor float64, destino *ContaCorrente) string {

	if valor <= c.saldo {
		if valor > 0 {
			c.saldo -= valor
			destino.Deposita(valor)
			return "Transferencia realizada com sucesso"
		} else if valor == 0 {
			return "Valor de transferência zerado"
		} else {
			return "Não é permitido valor negativo"
		}
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) MostraSaldo() (string, float64) {
	return "Saldo Atual: ", c.saldo
}
