package error

type Unwrapper interface {
  Unwrap() error
}
