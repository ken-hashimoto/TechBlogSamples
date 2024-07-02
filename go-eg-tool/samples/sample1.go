package samples

import (
	"errors"
	"fmt"
)

func sample1() {
	msg := "something went wrong"
	err := errors.New("error: " + msg)
	fmt.Println(err)
}
