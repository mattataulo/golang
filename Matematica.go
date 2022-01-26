package main

import "fmt"

//Função calcular contempla os casos de +, -, / e *.
func calcular(operacao string, numero1, numero2 int) (saida int) {
	switch operacao {
	case "+":
		saida = numero1 + numero2
	case "-":
		saida = numero1 - numero2
	case "*":
		saida = numero1 * numero2
	case "/":
		saida = numero1 / numero2
	default:
		fmt.Println("Operação Inválida, tente novamente.")
	}
	return saida
}

//Requisição dos numeros + apresentar os resultados
func main() {
	var operacao string
	var numero1, numero2 int
	fmt.Print("Insira o primeiro numero: ")
	fmt.Scanln(&numero1)
	fmt.Print("Insira o segundo numero: ")
	fmt.Scanln(&numero2)
	fmt.Print("Insira a operação desejada (+,-,/,*):")
	fmt.Scanln(&operacao)
	saida := calcular(operacao, numero1, numero2)

	fmt.Printf("%d %s %d = %d", numero1, operacao, numero2, saida)
}
