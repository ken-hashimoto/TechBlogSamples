package samples

import (
	"fmt"
	"time"
)

func sample3(){
	t := time.Now()
	isZero := t != time.Time{}
	if isZero{
		fmt.Println("now is not zero")
	}
}