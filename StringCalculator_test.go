package main

import "testing"

func TestStringCalculator(t *testing.T) {

	type args struct {
		numbers string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"Returns 0 when getting an empty string", args{""}, 0},
		{"Returns the value if only got one", args{"1"}, 1},
		{"Returns the addition of both numbers", args{"1,2"}, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.numbers); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
