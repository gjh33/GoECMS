package ecms

// Module represents atomic functions within your code. They should execute one simple task.
type Module interface {
	Perform()
}
