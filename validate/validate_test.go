package validate

import (
	"testing"
)

func TestValidateName_Error(t *testing.T) {
	err := ValidateName("")

	if err == nil {
		t.Fatal("ожидали ошибку, но получили nil")
	}

	if err != ErrEmptyName {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestValidateName_Success(t *testing.T) {
	err := ValidateName("Natasha")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
