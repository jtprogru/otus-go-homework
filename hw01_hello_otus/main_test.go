package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask0004Solution(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		whant string
	}{
		{
			name:  "Hello, OTUS!",
			s:     "Hello, OTUS!",
			whant: "!SUTO ,olleH",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := Reverse(tt.s)
			assert.Equal(t, res, tt.whant)
		})
	}
}
