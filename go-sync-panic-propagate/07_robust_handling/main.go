package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := robustTaskExecution(); err != nil {
		log.Printf("Task execution returned error: %v", err)
	}
}

func robustTaskExecution() error {
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case errgroup.PanicError:
				// error型のpanicは適切にログ出力して継続
				logError("Panic occurred", v.Recovered, string(v.Stack))
			case errgroup.PanicValue:
				// 不明な型のpanicは詳細情報と共にアラート
				sendAlert("Unknown panic type", v.Recovered, string(v.Stack))
			}
		}
	}()

	g := new(errgroup.Group)

	g.Go(riskyTask1)
	g.Go(riskyTask2)
	g.Go(riskyTask3)

	return g.Wait()
}

func riskyTask1() error {
	fmt.Println("Risk task 1: Starting")
	time.Sleep(100 * time.Millisecond)

	// 30%の確率でエラーを返す
	if rand.Float32() < 0.3 {
		return errors.New("task 1 failed with error")
	}

	// 10%の確率でpanicする
	if rand.Float32() < 0.1 {
		panic("task 1 panicked")
	}

	fmt.Println("Risk task 1: Completed successfully")
	return nil
}

func riskyTask2() error {
	fmt.Println("Risk task 2: Starting")
	time.Sleep(200 * time.Millisecond)

	// 20%の確率でerror型でpanic
	if rand.Float32() < 0.2 {
		panic(errors.New("task 2 panicked with error"))
	}

	// 5%の確率でstring型でpanic
	if rand.Float32() < 0.05 {
		panic("task 2 panicked with string")
	}

	fmt.Println("Risk task 2: Completed successfully")
	return nil
}

func riskyTask3() error {
	fmt.Println("Risk task 3: Starting")
	time.Sleep(150 * time.Millisecond)

	// 25%の確率でエラーを返す
	if rand.Float32() < 0.25 {
		return errors.New("task 3 failed with error")
	}

	fmt.Println("Risk task 3: Completed successfully")
	return nil
}

func logError(message string, err error, stack string) {
	fmt.Printf("ERROR LOG: %s - %v\n", message, err)
	fmt.Printf("Stack trace (first 200 chars): %s...\n", stack[:min(200, len(stack))])
}

func sendAlert(message string, value any, stack string) {
	fmt.Printf("ALERT: %s - %v\n", message, value)
	fmt.Printf("Stack trace (first 200 chars): %s...\n", stack[:min(200, len(stack))])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
