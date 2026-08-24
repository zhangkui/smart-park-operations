package workflows

import "context"

func ExportTenantIDs(ctx context.Context, tenants []string, write func(string) error) error {
	for _, tenant := range tenants {
		if err := write(tenant); err != nil {
			return err
		}
	}
	return nil
}
