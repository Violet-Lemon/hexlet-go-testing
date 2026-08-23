package currency

import "testing"

type mockConverter struct {
	lastAmount       float64
	lastFrom, lastTo string
	calls            int
}

func (c *mockConverter) Convert(amount float64, from, to string) float64 {
	c.lastAmount, c.lastFrom, c.lastTo = amount, from, to
	c.calls++

	return 42.0
}

func TestPriceIn_DelegatesAndReturns(t *testing.T) {
	mock := &mockConverter{}

	result := PriceIn(100.0, "usd", "euro", mock)

	if result != 42.0 {
		t.Errorf("ожидали 42.0, получили %v", result)
	}

	if mock.calls != 1 {
		t.Fatalf("calls: got %d, want 1", mock.calls)
	}

	if mock.lastAmount != 100 || mock.lastFrom != "usd" || mock.lastTo != "euro" {
		t.Fatalf("args mismatch: (%v,%s->%s)", mock.lastAmount, mock.lastFrom, mock.lastTo)
	}
}

func TestPriceIn_ZeroAndNegative(t *testing.T) {
	m := &mockConverter{}
	_ = PriceIn(0, "RUB", "USD", m)
	if m.lastAmount != 0 || m.lastFrom != "RUB" || m.lastTo != "USD" {
		t.Fatalf("args mismatch for zero: (%v,%s->%s)", m.lastAmount, m.lastFrom, m.lastTo)
	}
	_ = PriceIn(-7, "USD", "EUR", m)
	if m.lastAmount != -7 || m.lastFrom != "USD" || m.lastTo != "EUR" {
		t.Fatalf("args mismatch for negative: (%v,%s->%s)", m.lastAmount, m.lastFrom, m.lastTo)
	}
}
