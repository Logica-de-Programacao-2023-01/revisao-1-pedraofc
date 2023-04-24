package main

import (
	"fmt"
	"strings"
)

func countAverageWordLength(text string) (float64, error) {
	words := strings.Fields(text)
	if len(words) == 0 {
		return 0, fmt.Errorf("texto vazio")
	}

	var sum float64
	for _, word := range words {
		sum += float64(len(word))
	}
	return sum / float64(len(words)), nil
}

func main() {
	text := "O rato roeu a roupa do rei de Roma"
	avgLen, err := countAverageWordLength(text)
	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Média de letras por palavra: %.1f\n", avgLen)
	}
}
