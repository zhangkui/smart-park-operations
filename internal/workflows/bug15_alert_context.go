package workflows

import "context"

func PublishAlert(ctx context.Context, send func() error) error {
	return send()
}
