package apperrors

import "fmt"

type NotFound struct{}

type LimitAttemptReached struct{}

type InvalideURL struct{}

type Unexpected struct {
	Err error
}

func (n NotFound) Error() string {
	return "Shortened URL not found."
}

func (l LimitAttemptReached) Error() string {
	return "Limit Attempts Exceed. Try again in a few seconds."
}

func (u Unexpected) Error() string {
	return fmt.Errorf("unexpected error occorred: %w", u.Err).Error()
}

func (i InvalideURL) Error() string {
	return "Invalid URL."
}
