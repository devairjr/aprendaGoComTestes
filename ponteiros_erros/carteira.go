package ponteiroserros

import (
	"errors"
	"fmt"
)

type Carteira struct {
	saldo Bitcoin
}

type Bitcoin int

type Stringer interface {
	String() string
}

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Depositar(quantidade Bitcoin) {
	c.saldo += quantidade
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (c *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= quantidade
	return nil
}
