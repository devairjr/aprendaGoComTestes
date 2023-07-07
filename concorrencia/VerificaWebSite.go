package concorrencia

import "net/http"

func VerificaWebsite(url string) bool {
	resposta, err := http.Head(url)
	if err != nil {
		return false
	}

	if resposta.StatusCode != http.StatusOK {
		return false
	}

	return true
}
