package rate_limiter

import (
	"errors"
)

// The State value represents an allowed request proceeding result state.
type State int

const (
	// StateOk represents the request proceeding result is ok.
	StateOk State = iota
	// StateErr represents the request proceeding result is error, but not time out.
	StateErr
	// StateTimeout represents the request proceeding result is time out.
	StateTimeout
)

var (
	// The ErrLimitExceed is returned when the rate limiter is
	// triggered and the request is rejected due to limit exceeded.
	errLimitExceed = errors.New("rate limit exceeded")
)

// ErrLimitExceed produce an error that would be returned when the
// rate limiter is triggered and the request is rejected due to limit exceeding.
func ErrLimitExceed() error {
	return errLimitExceed
}

// IsErrLimitExceed checks whether an error is indicating whether
// a request is rejected due to limit exceeding.
// This function supports wrapped errors.
func IsErrLimitExceed(err error) bool {
	return errors.Is(err, errLimitExceed)
}

// DoneFunc is done function.
type DoneFunc func(DoneInfo)

// DoneInfo is done info.
type DoneInfo struct {
	State State
	Err   error
}

// Limiter is a rate limiter.
type Limiter interface {
	Allow() (DoneFunc, error)
}
