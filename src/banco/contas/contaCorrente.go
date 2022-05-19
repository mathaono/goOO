package contas

type ContaCorrente struct {
	Nome          string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Saque(valor float64) (string, float64) {
	if valor <= c.Saldo {
		if valor > 0 {
			c.Saldo -= valor
			return "Saque realizado com sucesso", valor
		} else {
			return "Não é permitido valor negativo", valor
		}
	} else {
		return "Saldo insuficiente", c.Saldo
	}
}

func (c *ContaCorrente) Deposita(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Deposito realizado com sucesso", valor
	} else {
		return "Impossível depositar valor negativo", valor
	}
}

func (c *ContaCorrente) Transfere(valor float64, destino *ContaCorrente) string {

	if valor <= c.Saldo {
		if valor > 0 {
			c.Saldo -= valor
			destino.Deposita(valor)
			return "Transferencia realizada com sucesso"
		} else if valor == 0 {
			return "Valor de transferência zerado"
		} else {
			return "Não é permitido valor negativo"
		}
	} else {
		return "Saldo insuficiente"
	}
}
