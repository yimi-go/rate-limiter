package rate_limiter

import (
	"errors"
)

// The State value represents an allowed request proceeding result state.
type State int

//goland:noinspection GoUnusedConst
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

// Token 请求许可凭证
type Token interface {
	// Ready 请求可执行信号
	Ready() <-chan struct{}
	// Done 请求完成回调
	Done(state State, err error)
}

// Limiter is a rate limiter.
type Limiter interface {
	// Allow checks all inbound traffic.
	// Once overload is detected, it raises limit.ErrLimitExceed error.
	//
	// If the request is permitted, the Token.Done method must be called after
	// finishing the request, even if the request is failed or timeout.
	Allow(id string) (Token, error)
}
