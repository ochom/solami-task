package tasks

import "testing"

func TestRunTask1(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case 1",
			args: args{1},
			want: "N/A",
		},
		{
			name: "case 2",
			args: args{4},
			want: "Fizz ",
		},
		{
			name: "case 3",
			args: args{10},
			want: "Fizz Buzz Fizz Fizz Buzz ",
		},
		{
			name: "case 4",
			args: args{15},
			want: "Fizz Buzz Fizz Fizz Buzz Fizz FizzBuzz ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunTask1(tt.args.N); got != tt.want {
				t.Errorf("RunTask1() = %v, want %v", got, tt.want)
			}
		})
	}
}
