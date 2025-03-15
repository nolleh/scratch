package error

import "testing"
import "errors"

func Test_Unwrap(t *testing.T) {
	t.Parallel()

	original := errors.New("original error")
	wrapped := Wrapper(original)

	unwrapped := errors.Unwrap(wrapped)
	if unwrapped != original {
		t.Fatalf("expected %v, got %v", original, unwrapped)
	}
}
