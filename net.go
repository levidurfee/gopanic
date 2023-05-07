package gopanic

import (
	"context"
	"errors"
)

// Copied from https://github.com/golang/go/blob/master/src/net/net.go#L424
func mapErr(err error) error {
	switch err {
	case context.Canceled:
		return errors.New("err")
	case context.DeadlineExceeded:
		return errors.New("err")
	default:
		return err
	}
}
