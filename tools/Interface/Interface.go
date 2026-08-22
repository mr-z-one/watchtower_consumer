package Interface

type IExecutable interface {
	Execute(args ...string) ([]byte, error)
}
