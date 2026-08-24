package workflows

func NotifyParkingChange(callback func(string), vehicle string) error {
	callback(vehicle)
	return nil
}
