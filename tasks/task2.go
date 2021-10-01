package tasks

//Eligibility is the loan eligibility status
type Eligibility string

const (
	//MIN_SALARY is the mimimum salary required for a user to qualify for a loan
	MIN_SALARY = 30000

	//MIN_WORK_YEARS minimum work years required to qualify for a loan
	MIN_WORK_YEARS = 2

	//eligibility status
	NOT_ELIGIBLE Eligibility = "Not Eligible"
	ELIGIBLE     Eligibility = "Eligible"
)

//CheckLoanEligibility checks if a client is eligible for a loan given

func CheckLoanEligibility(salary int, currentWorkDuration int, haveLoan bool) Eligibility {

	//client must earn this salary to qualify
	if salary < MIN_SALARY {
		return NOT_ELIGIBLE
	}
	if haveLoan {
		return NOT_ELIGIBLE
	}

	if currentWorkDuration < MIN_WORK_YEARS {
		return NOT_ELIGIBLE
	}

	return ELIGIBLE

}
