package centipede

import (
	"context"
	"errors"
)

var (
	ErrExecutionCanceled error = errors.New("execution canceled")
)

// RunWithContext accepts a context and a function that produces T
// The function will be run, and so long as the context is not done,
// its result will be returned. Otherwise an error will be returned.
func RunWithContext[T any](ctx context.Context, f func() T) (*T, error) {
	ch := make(chan T, 1)
	go func() {
		r := f()
		ch <- r
	}()

	select {
	case res := <-ch:
		return &res, nil
	case <-ctx.Done():
		return nil, ErrExecutionCanceled
	}
}
