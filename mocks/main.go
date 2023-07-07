package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Pausa()
}

type SleeperConfiguravel struct {
	duracao time.Duration
	pausa   func(time.Duration)
}

func (s *SleeperConfiguravel) Pausa() {
	s.pausa(s.duracao)
}

const ultimaPalavra = "Vai!"
const inicioContagem = 3

func Contagem(saida io.Writer, sleeper Sleeper) {
	for i := inicioContagem; i > 0; i-- {
		sleeper.Pausa()
		fmt.Fprintln(saida, i)
	}
	sleeper.Pausa()
	fmt.Fprint(saida, ultimaPalavra)
}

func main() {
	sleeper := &SleeperConfiguravel{1 * time.Second, time.Sleep}
	Contagem(os.Stdout, sleeper)
}
