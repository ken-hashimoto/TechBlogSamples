package main

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

// v0.14.0以降（panic が Wait() から伝播されることを期待）
func TestPanicBehavior(t *testing.T) {
	g := new(errgroup.Group)
	g.Go(func() error {
		panic("test panic")
		return nil
	})

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic from Wait()")
		} else {
			t.Logf("Successfully caught panic: %v", r)
		}
	}()

	g.Wait()
}

// PanicError型のテスト
func TestPanicErrorType(t *testing.T) {
	g := new(errgroup.Group)
	g.Go(func() error {
		panic(errgroup.PanicError{})
		return nil
	})

	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(errgroup.PanicError); ok {
				t.Log("Successfully caught PanicError")
			} else {
				t.Errorf("Expected PanicError, got %T", r)
			}
		} else {
			t.Error("Expected panic")
		}
	}()

	g.Wait()
}

// PanicValue型のテスト
func TestPanicValueType(t *testing.T) {
	g := new(errgroup.Group)
	g.Go(func() error {
		panic("string panic")
		return nil
	})

	defer func() {
		if r := recover(); r != nil {
			if pv, ok := r.(errgroup.PanicValue); ok {
				t.Logf("Successfully caught PanicValue: %v", pv.Recovered)
			} else {
				t.Errorf("Expected PanicValue, got %T", r)
			}
		} else {
			t.Error("Expected panic")
		}
	}()

	g.Wait()
}
