package workflows

// ParkingNotifier delivers optional gate notifications.
type ParkingNotifier interface {
	Notify(vehicleID, status string)
}

func NotifyParkingGate(notifier ParkingNotifier, vehicleID, status string) {
	notifier.Notify(vehicleID, status)
}
