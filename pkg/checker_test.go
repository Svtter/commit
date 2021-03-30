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
			if !CheckPrefix(tt.msg) == tt.valid {
				t.Errorf("tt.msg is not equal to tt.valid")
			}
		})
	}
}