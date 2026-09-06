package testing

import "testing"

func TestCalculateTotal(t *testing.T) {

	products := []Product{
		{Name: "A", Price: 100},
		{Name: "B", Price: 200},
	}

	total := CalculateTotal(products)

	if total != 300 {
		t.Errorf("expected 300 got %d", total)
	}
}
