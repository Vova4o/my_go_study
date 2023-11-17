//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"fmt"
	"time"
)

func TimeParser(str string) (time.Time, error) {
	var new time.Time

	var err error

	new, err = time.Parse(time.TimeOnly, str)
	fmt.Println("error", err)
	if err != nil {
		return time.Time{}, fmt.Errorf("string %v, is not time. error: %v", str, err)
	} else {
		return new, nil
	}
}
