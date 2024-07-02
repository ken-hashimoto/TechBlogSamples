package template1

import (
	"errors"
	"fmt"
)

// eg -t ./templates/template1/template1.go -w ./samples/sample1.go

func before(s string) error { return fmt.Errorf("%s", s) }
func after(s string)  error { return errors.New(s) }