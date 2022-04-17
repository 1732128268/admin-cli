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
				src: "123456",
			},
			want: "313233343536d41d8cd98f00b204e9800998ecf8427e",
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
