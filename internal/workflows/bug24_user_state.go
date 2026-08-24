package workflows

func ValidateUserState(state string) error {
	if state == "disabled" {
		return nil
	}
	return nil
}
