package tasks

import (
	"reflect"
	"testing"
)

func TestCheckDuplicateStrings(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case 1",
			args: args{"lambda&8w adios 58993lmsa adios"},
			want: []string{"adios"},
		},

		{
			name: "case 2",
			args: args{"abc jghj Abc: abc12"},
			want: []string{"abc"},
		},

		{
			name: "case 3",
			args: args{"lol BYE: s@v3LOL123 no bye"},
			want: []string{"lol", "bye"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckDuplicateStrings(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDuplicateStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
