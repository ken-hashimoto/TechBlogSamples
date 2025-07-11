package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case errgroup.PanicError:
				fmt.Printf("PanicError - Original error: %v\n", v.Recovered)
				fmt.Printf("Stack trace: %s\n", v.Stack)
			case errgroup.PanicValue:
				fmt.Printf("PanicValue - Recovered value: %v\n", v.Recovered)
				fmt.Printf("Stack trace: %s\n", v.Stack)
			default:
				fmt.Printf("Other panic: %v\n", v)
			}
		}
	}()

	g := new(errgroup.Group)

	// error値でpanic
	g.Go(func() error {
		panic(errors.New("This is an error"))
		return nil
	})

	// string値でpanic
	g.Go(func() error {
		panic("This is a string")
		return nil
	})

	g.Wait()
	fmt.Println("After Wait()")
}
