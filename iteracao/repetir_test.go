package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a", 10)
	esperado := "aaaaaaaaaa"

	if repeticoes != esperado {
		t.Errorf("esperado '%s' mas obteve '%s'", esperado, repeticoes)
	}

}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 10)
	}
}

func ExampleRepetir() {
	repeticoes := Repetir("a", 10)
	fmt.Println(repeticoes)
	// Output: aaaaaaaaaa
}
