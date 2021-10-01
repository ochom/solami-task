package main

import (
	"fmt"

	"github.com/ochom/solami/interview/tasks"
	"github.com/sirupsen/logrus"
)

func main() {
	//task 1
	fmt.Println(tasks.RunTask1(15))

	//task 2
	fmt.Println(tasks.CheckLoanEligibility(35000, 3, true))

	//task 3
	logrus.Println(tasks.CheckDuplicateStrings("lambda&8w adios 58993lmsa adios"))
}
