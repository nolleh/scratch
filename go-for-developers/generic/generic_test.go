package generic

import "slices"

import "fmt"
import "testing"

func Test_Keys(t *testing.T) {
	t.Parallel()

	m := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	act := Keys(m)

	slices.Sort(act)

	exp := []int{1, 2, 3}

	al := len(act)
	el := len(exp)
	if el != al {
		t.Fatalf("expected len %d, but got %d", el, al)
	}

	at := fmt.Sprintf("%T", act)
	et := fmt.Sprintf("%T", exp)

	if at != et {
		t.Fatalf("expected type %s, but got type %s", et, at)
	}

}
