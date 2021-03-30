package pkg

import (
	"testing"
)

func TestCheckPrefix(t *testing.T) {
	tests := []struct {
		name  string
		msg   string
		valid bool
	}{
		{
			"failed case test",
			"hello",
			false,
		},
		{
			"success case test",
			"fix: problems",
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPrefix(tt.msg)
			if got != tt.valid {
				t.Errorf("CheckPrefix() = %v, want %v", got, tt.valid)
			}
		})
	}
}