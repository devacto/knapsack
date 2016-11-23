package main

import "testing"

func Test_throw(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test 1", want: "a string"},
		{name: "test 2", want: "a string"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := throw(); got != tt.want {
				t.Errorf("throw() = %v, want %v", got, tt.want)
			}
		})
	}
}
