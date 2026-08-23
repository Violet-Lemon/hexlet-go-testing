package greet

import "testing"

func TestHello(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		got, err := Hello("Go")
		assertError(t, err, false)
		assertString(t, got, "Hello, Go")

		got, err = Hello("World")
		assertError(t, err, false)
		assertString(t, got, "Hello, World")
	})

	t.Run("empty", func(t *testing.T) {
		got, err := Hello("")
		assertError(t, err, true)
		assertString(t, got, "")
	})

	t.Run("unicode", func(t *testing.T) {
		got, err := Hello("Гофер")
		assertError(t, err, false)
		assertString(t, got, "Hello, Гофер")
	})

	t.Run("trim", func(t *testing.T) {
		got, err := Hello(" Go ")
		assertError(t, err, false)
		assertString(t, got, "Hello,  Go ")
	})
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("получили строку %q, хотели %q", got, want)
	}
}

func assertError(t testing.TB, got error, want bool) {
	t.Helper()
	if (got != nil) != want {
		if want {
			t.Fatalf("ожидалась ошибка, но её нет")
		}
		t.Fatalf("не ожидалась ошибка, но получена: %v", got)
	}
}
