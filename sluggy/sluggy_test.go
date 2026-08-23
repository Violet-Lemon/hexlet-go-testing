package sluggy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlug(t *testing.T) {
	tests := []struct {
		name, in, want string
	}{
		{"punctuation", "Best Idea!", "best-idea"},
		{"dublicate separator", "sunny____weather!", "sunny-weather"},
		{"unicode", "солнце Чудесное", "солнце-чудесное"},
		{"empty", "", ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Slug(test.in)
			assert.Equal(t, test.want, got)
		})
	}
}
