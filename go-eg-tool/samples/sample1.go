package samples

import (
	"fmt"
)

func sample1() {
    msg := "something went wrong"
    err := fmt.Errorf("%s", "error: " + msg)
    fmt.Println(err)
}