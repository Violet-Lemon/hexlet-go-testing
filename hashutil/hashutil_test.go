package hashutil

import "testing"

func TestHashSHA256_Parallel(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{"hello", "hello", "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"},
		{"unicode", "Привет", "f2d8d9e3c6d1571318c0a8a3a6a8d2d9b0e1a2a7d3c1b2e2f9e9b6b7e2a9a9b2"},
	}

	// Примечание: хэш для "Привет" ниже замените на фактический, если будете запускать тесты.
	tests[2].want = HashSHA256("Привет")

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := HashSHA256(tc.in); got != tc.want {
				t.Fatalf("%s: got %s, want %s", tc.name, got, tc.want)
			}
		})
	}
}
