package template6

import (
	"fmt"
	"strconv"
)

// eg -t ./template.go -w ./samples/sample6.go

func before(x int) { fmt.Println(x) }
func after(x int) { fmt.Println(strconv.Itoa(x)) }