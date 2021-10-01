package tasks

//RunTask1 computes numbers between 1 and 100 inclusive and prints
//`Fizz` for multiples of 3
//`Buzz` for multiples of 5
//`FizzBuzz` for multiples of both 3 n 5
//`N/A`  otherwise
func RunTask1(N int) string {
	numbersFit := false
	for i := 1; i <= N; i++ {
		if !numbersFit {
			if i%3 == 0 || i%5 == 0 {
				numbersFit = true
			}
		}
	}

	//if no number fits
	if !numbersFit {
		return "N/A"
	} else {
		str := ""
		for i := 1; i <= N; i++ {
			if i%3 == 0 && i%5 == 0 {
				str += "FizzBuzz "
			} else if i%3 == 0 {
				str += "Fizz "
			} else if i%5 == 0 {
				str += "Buzz "
			}
		}
		return str
	}
}
