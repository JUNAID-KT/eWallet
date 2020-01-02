package util

import (
	"math/rand"
	"time"

	"github.com/JUNAID-KT/eWallet/models"
)

func SetStatus(statusCode int, descriptionCode string, description string) models.Status {
	return models.Status{
		Status: models.StatusResponse{
			StatusCode:      statusCode,
			DescriptionCode: descriptionCode,
			Description:     description,
		},
	}
}
func Retry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if attempts--; attempts > 0 {
			// Add some randomness to prevent creating race condition
			jitter := time.Duration(rand.Int63n(int64(sleep)))
			sleep = sleep + jitter/2
			time.Sleep(sleep)
			return Retry(attempts, 2*sleep, f)
		}
		return err
	}
	return nil
}
