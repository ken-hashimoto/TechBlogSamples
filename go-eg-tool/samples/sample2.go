package samples

import "io/ioutil"

func sample2() {
	_,_ = ioutil.ReadFile("filename")
}

// eg -t ./templates/template2/template2.go -w .samples/sample2.go