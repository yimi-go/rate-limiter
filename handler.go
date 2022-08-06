package rate_limiter

import (
	"context"
)

// Handler only handles the execution when it's rate limiter allows.
type Handler interface {
	// Handle executes the execFn if the rate limiter allows.
	// The execFn returns any error occurred during executing.
	// The Handle function returns a bool indicating whether the execFn is executed.
	// A caller can also recording execution in the execFn ignoring the returned bool.
	Handle(ctx context.Context, execFn func() error) (executed bool)
}

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as limit handlers. If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler that calls f.
type HandlerFunc func(ctx context.Context, execFn func() error) (executed bool)

func (h HandlerFunc) Handle(ctx context.Context, execFn func() error) (executed bool) {
	return h(ctx, execFn)
}
