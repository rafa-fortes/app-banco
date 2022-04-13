package contas //modularizamos nosso código criando um pacote para manter os diferentes tipos de conta

import "banco/clientes"

type ContaCorrente struct { // Para criar uma nova struct, usamos o prefixo type, seguido do nome da struct, o sufixo struct
	Titular                       clientes.Titular // e declaramos os campos entre chaves, conforme o exemplo abaixo:
	NumeroAgencia, NumeroCorrente int
	saldo                         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string { //Função para verificar o saldo da conta, chamado Sacar.
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saca realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

//Aqui estmos apontando para conta corrente que está chamando. c *ContaCorrente
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) { //Função para fazer o deposito na conta
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso. saldo atual:", c.saldo // retornando o valor após o deposito
	} else {
		return "Valor do deposito menor que zero. saldo atual ", c.saldo
	}
}

//Como queremos alterar o valor dessa conta também, colocamos um asterisco na frente dela nos parâmetros da função.
func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool { // Função para a transferencia de dinheiro.
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 { // Verificando se existe saldo na conta e verificando se o valor da transferencia é > 0
		c.saldo -= valorDaTransferencia              // Retirando o valor da conta para a transferencia para a outra conta.
		contaDestino.Depositar(valorDaTransferencia) // Transferindo o dinheiro para a conta, usando a outra função "Depositar".
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 { //Função para ver o saldo da conta.
	return c.saldo
}
