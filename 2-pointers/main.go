package main

type Conta struct {
	saldo int
}

func (c Conta) simulador(valor int) int {
	c.saldo += valor
	println(c.saldo)
	return valor
}

func main() {

	conta := Conta{
		saldo: 100,
	}
	conta.simulador(200)
	println(conta.saldo)
}
