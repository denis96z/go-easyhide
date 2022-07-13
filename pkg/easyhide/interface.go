package easyhide

type Marshaler interface {
	EasyHide() ([]byte, error)
}
