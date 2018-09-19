package compile

import (
	"fmt"
)

var TestFailedError = fmt.Errorf("Some test cases did not pass")

func RunTests(args ...string) error {
	if len(args)%3 != 0 {
		fmt.Println("Invalid test setup, expects args to be a multiple of 3")
		return fmt.Errorf("Invalid test setup")
	}

	errors := make([]error, 0)
	erroringCases := make([]string, 0)
	for i := 0; i < len(args); i += 3 {
		err := Assert(args[i], args[i+1], args[i+2])
		if err != nil {
			errors = append(errors, err)
			erroringCases = append(erroringCases, args[i+2])
		}
	}

	if len(errors) == 0 {
		return nil
	}
	for i, v := range errors {
		fmt.Println(erroringCases[i], ":")
		fmt.Println("\tΔ", v.Error())
	}
	return TestFailedError
}
