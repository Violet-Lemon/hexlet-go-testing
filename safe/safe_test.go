package safe

import "testing"

func TestMustAt_OutOfrange(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику, но её не было")
		} else if r != "index out of range" {
			t.Errorf("неожиданное сообщение паники: %v", r)
		}
	}()

	_ = MustAt([]int{1, 2, 3}, 7)
}

func TestMustAt_NegativeIndex(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("ожидали панику, но её не было")
		} else if r != "index out of range" {
			t.Errorf("неожиданное сообщение паники: %v", r)
		}
	}()

	_ = MustAt([]int{1, 2, 3}, -4)
}

func TestMustAt_Valid(t *testing.T) {
	xs := []int{10, 20, 30}
	got := MustAt(xs, 1)
	if got != 20 {
		t.Fatalf("got %d, want %d", got, 20)
	}
}
