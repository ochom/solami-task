package tasks

import "testing"

func TestCheckLoanEligibility(t *testing.T) {
	type args struct {
		salary              int
		currentWorkDuration int
		haveLoan            bool
	}
	tests := []struct {
		name string
		args args
		want Eligibility
	}{
		{
			name: "case 1",
			args: args{35000, 3, true},
			want: NOT_ELIGIBLE,
		},
		{
			name: "case 2",
			args: args{28000, 4, false},
			want: NOT_ELIGIBLE,
		},
		{
			name: "case 3",
			args: args{31000, 6, false},
			want: ELIGIBLE,
		},
		{
			name: "case 4",
			args: args{32000, 1, false},
			want: NOT_ELIGIBLE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckLoanEligibility(tt.args.salary, tt.args.currentWorkDuration, tt.args.haveLoan); got != tt.want {
				t.Errorf("CheckLoanEligibility() = %v, want %v", got, tt.want)
			}
		})
	}
}
