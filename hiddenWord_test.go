package main

import "testing"

func TestHiddenWord(t *testing.T) {

	type args struct {
		key uint32
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"Key: 147", args{147}, "bed"},
		{"Key: 637", args{637}, "aid"},
		{"Key: 7415", args{7415}, "debt"},
		{"Key: 49632", args{49632}, "email"},
		{"Key: 942547", args{942547}, "melted"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HiddenWord(tt.args.key); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
