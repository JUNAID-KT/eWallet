package search_engine

import (
	"context"
	"errors"
	"net/http"
	"syscall"
	"time"

	"github.com/JUNAID-KT/eWallet/util"

	"github.com/olivere/elastic"
)

type RetryPolicy struct {
	backOff elastic.Backoff
}

func CustomRetrier() *RetryPolicy {
	return &RetryPolicy{
		backOff: elastic.NewExponentialBackoff(10*time.Millisecond, 8*time.Second),
	}
}

func (r *RetryPolicy) Retry(ctx context.Context, retry int, req *http.Request, resp *http.Response, err error) (time.Duration, bool, error) {
	// Fail hard on a specific error
	if err == syscall.ECONNREFUSED {
		return 0, false, errors.New("elastic search or network down")
	}

	// Stop after 5 retries
	if retry >= util.MaxRetries {
		return 0, false, nil
	}

	// Let the backOff strategy decide how long to wait and whether to stop
	wait, stop := r.backOff.Next(retry)
	return wait, stop, nil
}
