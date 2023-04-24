package main

import (
	"errors"
	"fmt"
)

type compra struct {
	valor float64
}

type historicoCompras []compra

func (h historicoCompras) valorTotal() float64 {
	total := 0.0
	for _, c := range h {
		total += c.valor
	}
	return total
}

func calcularDesconto(valorCompra float64, historico historicoCompras) (float64, error) {
	if valorCompra <= 0 {
		return 0, errors.New("valor da compra inválido")
	}

	var desconto float64

	if len(historico) == 0 {
		// primeira compra do cliente
		desconto = valorCompra * 0.1
	} else {
		totalCompras := historico.valorTotal()
		mediaCompras := totalCompras / float64(len(historico))

		if totalCompras > 1000 {
			desconto = valorCompra * 0.1
		} else if totalCompras <= 500 {
			desconto = valorCompra * 0.02
		} else {
			desconto = valorCompra * 0.05
		}

		if mediaCompras > 1000 && desconto < valorCompra*0.2 {
			desconto = valorCompra * 0.2
		}
	}

	return desconto, nil
}

func main() {
	historico := historicoCompras{{valor: 500}, {valor: 750}}
	desconto, err := calcularDesconto(1000, historico)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Desconto:", desconto)
	}
}
