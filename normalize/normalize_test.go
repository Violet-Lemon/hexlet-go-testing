package normalize

import "testing"

func TestClean(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"     a  b	\t\t\t		c   ", "a b c"},
		{"", ""},
		{"AndHHHj", "andhhhj"},
		{"kkklaaa", "kkklaaa"},
		{"   home is So LLLovely :*", "home is so lllovely :*"},
	}

	for _, c := range cases {
		got := Clean(c.in)

		if got != c.want {
			t.Errorf("Результат %s, ожидали %s", got, c.want)
		}
	}
}
