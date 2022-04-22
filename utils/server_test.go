package utils

import "testing"

func TestRound(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "TestRound",
			args: args{
				f: 1.55555,
			},
			want: 1.55,
		},
		{
			name: "TestRound",
			args: args{
				f: 1.4,
			},
			want: 1.4,
		},
		{
			name: "TestRound",
			args: args{
				f: 1.6,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Round(tt.args.f); got != tt.want {
				t.Errorf("Round() = %v, want %v", got, tt.want)
			}
		})
	}
}
