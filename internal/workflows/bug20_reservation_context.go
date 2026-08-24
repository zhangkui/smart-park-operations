package workflows

import "context"

func ReserveSpace(ctx context.Context, reserve func() error) error {
	return reserve()
}
