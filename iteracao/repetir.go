package iteracao

import "strings"

func Repetir(caractere string, rep int) string {
	/*var repeticoes string

	for i := 0; i < rep; i++ {
		repeticoes += caractere
	}
	return repeticoes*/

	return strings.Repeat(caractere, rep)
}
