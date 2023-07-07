package arrays_slices

func Soma(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func SomaTudo(numerosParaSomar ...[]int) (somas []int) {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas = make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma(numeros)
	}

	return
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma(final))
		}
	}

	return somas
}
