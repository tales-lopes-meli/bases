package main

import (
	print "github.com/tales-lopes-meli/bases/print"
	climate "github.com/tales-lopes-meli/bases/climate"
)

func main() {
	print.PrintInfo("Tales", 22)
	climate.ShowClimate(1, 2, 3)

	// EXERCICIO 3
	// var 1nome string -> Precisa começar com letra
	// var sobrenome string -> Certo
	// var int idade -> O tipo de dados vem após o nome da variável
	// 1sobrenome := 6 -> Precisa começar com letra
	// var licenca_para_dirigir = true -> Deveria usar o camelCase
	// var estatura da pessoa int -> Utilizar espaços não é permitido
	// quantidadeDeFilhos := 2 -> Certo

	// EXERCICIO 4
	// var sobrenome string = "Silva" -> Certo
	// var idade int = "25" -> Tipo diferente, retirar as aspas
	// boolean := "false"; -> Os valores do tipo bool não necessitam de aspas
	// var salario string = 4585.90 -> Está errado, deveria usar o tipo float32 ou colocar aspas no valor
	// var nome string = "Fellipe" -> Certo


}