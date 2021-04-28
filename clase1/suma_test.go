package clase1

import "testing"

func TestAdd(t *testing.T) {
	want := 5
	t.Logf("want vale: %d\n", want)
	got := Add(2, 3)
	t.Logf("got vale: %d\n", got)
	if got != want {
		//t.Logf("Error: Se esperaba %d, se obtuvo %d", want, got)
		//t.Fail()
		t.Errorf("Error: Se esperaba %d, se obtuvo %d", want, got)
	}
	t.Log("Termino la prueba de Add")
}

func TestAddAcum(t *testing.T) {
	want := 6
	got := AddAcum(1, 2, 3)

	if got != want {
		//t.Logf("Error: Se esperaba %d, se obtuvo %d", want, got)
		//t.Fail()
		t.Errorf("Error: Se esperaba %d, se obtuvo %d", want, got)
	}
}
