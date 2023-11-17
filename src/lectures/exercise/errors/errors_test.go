package timeparse

import (
	"fmt"
	"testing"
	"time"
)

var tCheck = time.Date(1900, 01, 01, 13, 34, 30, 0, time.UTC)

// fmt.Println(tCheck)
// fmt.Printf(tCheck.Format("15:04:05"))

func TestTimeParser(t *testing.T) {
	res, err := TimeParser("text")
	if res.Format("15:04:05") != tCheck.Format("15:04:05") {
		fmt.Println(err)
	}

}
