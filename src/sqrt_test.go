package main
import "testing"

func TestSqrt( t *testing.T ) {

	deveRetornarVinteNove := sqrt(16, 2)
	
	if deveRetornarVinteNove != 29.41906792896077 {
		t.Errorf("deveRetornarVinteNove teste FALHOU: esperado o valor %v, mas retornou %v", 2, deveRetornarVinteNove)
	} else {
		t.Logf("deveRetornarVinteNove teste OK")
	}
}