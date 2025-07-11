package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := new(errgroup.Group)

	g.Go(func() error {
		fmt.Println("Normal task completed")
		return nil
	})

	g.Go(func() error {
		panic("Something went wrong!")
		return nil
	})

	g.Go(func() error {
		panic(errors.New("Error wrapped in panic"))
		return nil
	})

	// Wait()でpanicが捕捉される
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered panic: %v\n", r)
		}
	}()

	if err := g.Wait(); err != nil {
		fmt.Printf("Error from Wait(): %v\n", err)
	}

	fmt.Println("Program continues to run!")
}
