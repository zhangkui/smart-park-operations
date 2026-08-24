package workflows

import "strings"

func MatchPermission(rolePermissions []string, required string) bool {
	for _, permission := range rolePermissions {
		if strings.HasPrefix(required, permission) {
			return true
		}
	}
	return false
}
