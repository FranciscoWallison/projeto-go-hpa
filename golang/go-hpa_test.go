package main
import (
  "testing"
)
//somando a raiz quadrada
func TestGoHpa(t *testing.T){
	got:= executa_calculo()
	want:= 2499999414254808.5

	if got != want {
		t.Errorf("errado")
	}
}