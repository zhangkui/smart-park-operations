package workflows

func ApplyWorkOrderTransition(lifecycle *WorkOrderLifecycle, id, state string) error {
	lifecycle.Move(id, state, "system", "workflow")
	return nil
}
