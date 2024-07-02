package main

/*
template1
eg -t ./template.go -w ./samples/sample1.go
*/
// func before(s string) error { return fmt.Errorf("%s", s) }
// func after(s string)  error { return errors.New(s) }

/*
template2
eg -t ./template.go -w ./samples/sample2.go
*/
// func before(filename string) ([]byte, error) {
//     return ioutil.ReadFile(filename)
// }
// func after(filename string) ([]byte, error) {
//     return os.ReadFile(filename)
// }

/*
template3
eg -t ./template.go -w ./samples/sample3.go
*/
// func before(t time.Time) bool {
//     return t != time.Time{}
// }
// func after(t time.Time) bool {
//     return !t.IsZero()
// }

/*
template4
eg -t ./template.go -w ./samples/sample4.go
*/
// func before(x int) int { return x + x }
// func after(x int) int { return 2 * x }

/*
template5
eg -t ./template.go -w ./samples/sample5.go
*/
// func before(x int) int { return x + x }
// func after(x int) int { return 2 * x }

/*
template6
eg -t ./template.go -w ./samples/sample6.go
*/
// func before(x int) { fmt.Println(x) }
// func after(x int) { fmt.Println(strconv.Itoa(x)) }