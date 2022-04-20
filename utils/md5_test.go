package utils

import "testing"

func TestGenMd5(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				src: "111111",
			},
			want: "96e79218965eb72c92a549dd5a330112",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenMd5(tt.args.src); got != tt.want {
				t.Errorf("GenMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}
