package main

import (
	"fmt"
	"runtime"

	"golang.org/x/sync/errgroup"
)

func main() {
	defer func() {
		fmt.Println("Main function is finishing")
	}()

	g := new(errgroup.Group)

	g.Go(func() error {
		fmt.Println("Task 1 completed")
		return nil
	})

	g.Go(func() error {
		fmt.Println("Task 2 calling Goexit")
		runtime.Goexit() // goroutineを強制終了
		return nil
	})

	fmt.Println("Before Wait()")
	g.Wait()
	fmt.Println("After Wait() - this line is never reached")
}
