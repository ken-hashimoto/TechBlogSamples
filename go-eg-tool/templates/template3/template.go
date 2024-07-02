package template3

import "time"

// eg -t ./templates/template3/template.go -w ./samples/sample3.go

func before(t time.Time) bool {
    return t != time.Time{}
}
func after(t time.Time) bool {
    return !t.IsZero()
}