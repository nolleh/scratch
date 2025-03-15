package error

func Wrapper(original error) error {
	original = ErrorB{err: original}
	original = ErrorC{err: original}
	original = ErrorA{err: original}
	return original
}

func WrapperLong(original error) error {
	return ErrorC{
		err: ErrorB{
			err: ErrorA{
				err: original,
			},
		},
	}
}
