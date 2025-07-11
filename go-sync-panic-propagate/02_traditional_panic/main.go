package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	g := new(errgroup.Group)

	g.Go(func() error {
		fmt.Println("Normal task completed")
		return nil
	})

	g.Go(func() error {
		panic("Something went wrong!")
		return nil
	})

	// この行に到達する前にプログラムが停止してしまう（v0.13.0以前）
	// v0.14.0以降ではpanicが伝播される
	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("This line is now reached with v0.14.0!")
}
