package extision

import "testing"

func TestUnitAdd(t *testing.T) {
	sum := UnitAdd()
	if sum != 55 {
		t.Fatal("fail")
	}
	t.Log("sucess")

}

func TestSub(t *testing.T) {
	sub := sub(10, 2)
	if sub != 8 {
		t.Fatal("fail")
	}
	t.Log("sucess")
}
