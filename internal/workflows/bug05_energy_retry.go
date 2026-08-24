package workflows

import (
	"context"
	"time"
)

// RetryEnergyAggregation retries a failed meter aggregation job.
func RetryEnergyAggregation(ctx context.Context, attempts int, retryDelay time.Duration, aggregate func() error) error {
	var err error
	for attempt := 0; attempt < attempts; attempt++ {
		err = aggregate()
		if err == nil {
			return nil
		}
		time.Sleep(retryDelay)
	}
	return err
}
