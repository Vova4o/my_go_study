package display

import "fmt"

// Ony Capital named functions will be imported
// if you name it with lower letters it will be only
// in this pacckage and wont be able to use
// outside of this package
func Display(msg string) {
	fmt.Println(msg)
}
