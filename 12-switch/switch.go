package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Valor incorreto"
	}
}

func diaDaSemana2(numero int) string {
	var diaSemana string
	switch {
		case numero == 1:
			diaSemana = "Domingo"
			fallthrough // joga para a próxima condição (uso muito raro)
		case numero == 2:
			diaSemana = "Segunda-Feira"
		case numero == 3:
			diaSemana = "Terça-Feira"
		case numero == 4:
			diaSemana = "Quarta-Feira"
		case numero == 5:
			diaSemana = "Quinta-Feira"
		case numero == 6:
			diaSemana = "Sexta-Feira"
		case numero == 7:
			diaSemana = "Sábado"
		default:
			diaSemana = "Valor incorreto"
	}
	return diaSemana
}

func main() {
	dia := diaDaSemana(5)
	fmt.Println(dia)
}