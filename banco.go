package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

// func main() { // Aqui temos dois tipos para usar a struct, usaremos sempre o segundo modo, pois com ele não precisaremos escrever o nome de todos os campos novamente.
// contaDoRafael := ContaCorrente{titular: "Rafael", numeroAgencia: 589, numeroCorrente: 123456, saldo: 125.5}

// contaDaFernada := ContaCorrente{"Fernanda", 222, 111222, 2000} // Se não precisarmos de todos os campos, por exemplo, apenas titular e saldo, teríamos que escrevê-los para especificar que os valores se tratam dos dois.

// fmt.Println(contaDoRafael)
// fmt.Println(contaDaFernada)

// contaDaSilvia := ContaCorrente{}
// contaDaSilvia.titular = "Silvia"
// contaDaSilvia.saldo = 500

// fmt.Println("Saldo em conta:", contaDaSilvia.saldo)
// fmt.Println(contaDaSilvia.Sacar(200)) // Chamando a função Sacar e passando o valor do saque
// fmt.Println(contaDaSilvia.Depositar(500)) //Chamando a função Depositar e depositando o valor.

// contaDaSilvia := contas.ContaCorrente{Titular: "Silva", Saldo: 300}
// contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

// status := contaDaSilvia.Transferir(-200, &contaDoGustavo) //Chamando a função Transferir e passando o valor do deposito.
//Precisamos identificar contaDoGustavo. Conseguimos fazer isso colocando um &
//quando fizermos a transferência. Assim, dizemos que queremos transferir de fato
//para esse endereço além de termos apontado para a conta.

// fmt.Println(status)
// fmt.Println(contaDaSilvia)
// fmt.Println(contaDoGustavo)

func main() {
	// clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor Go"}
	// contaDoBruno := contas.ContaCorrente{clienteBruno, 123, 123456, 100}
	// fmt.Println(contaDoBruno)
	contaDoDenis := contas.ContaPoupanca{} //Criando as contas
	contaDoDenis.Depositar(100)
	PagarBoleto(&contaDoDenis, 60)
	//contaDoDenis.Sacar(5000)

	fmt.Println("Saldo atual:", contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(500)
	PagarBoleto(&contaDaLuiza, 400) // para pagar o boleto porque a conta não implementa o tipo verificarConta, pois nosso método Sacar() tem um receptor,ou seja, existe um endereço que ele busca, pois fizemos (c * ContaCorrente) antes da função, e nós não o passamos. Para passá-lo, precisaremos colocar o símbolo de "&" na frente de contaDoDenis

	fmt.Println("Saldo atual:", contaDaLuiza.ObterSaldo())
}
