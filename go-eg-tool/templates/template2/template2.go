package template2

import (
	"io/ioutil"
	"os"
)

// eg -t ./templates/template2/template2.go -w ./samples/sample2.go

func before(filename string) ([]byte, error) {
    return ioutil.ReadFile(filename)
}
func after(filename string) ([]byte, error) {
    return os.ReadFile(filename)
}