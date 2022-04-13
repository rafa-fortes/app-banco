package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string { //Função para verificar o saldo da conta, chamado Sacar.
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saca realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//Aqui estmos apontando para conta corrente que está chamando. c *ContaCorrente
func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) { //Função para fazer o deposito na conta
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. saldo atual:", c.saldo // retornando o valor após o deposito
	} else {
		return "Valor do deposito menor que zero. saldo atual ", c.saldo
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 { //Função para ver o saldo da conta.
	return c.saldo
}
