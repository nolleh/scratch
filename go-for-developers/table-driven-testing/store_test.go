package store

import "testing"
import "errors"


func noData(t testing.TB) *Store {
  return &Store{}
}

func withData(t testing.TB) *Store {
  return &Store{
    data: data{},
  }
}

func withUsers(t testing.TB) *Store {
  t.Helper()
  t.Fatal("not implemented")
  return nil
}


func Test_Store_All_Erros(t *testing.T) {
	t.Parallel()

	tn := "users"
	noData := &Store{}
	withData := &Store{data: data{}}
	// withUsers := &Store{data: data{"users": Models{}}}
  withUsers := withUsers(t)

	tcs := []struct {
		name  string
		store *Store
		exp   error
	}{
		{name: "nodata", store: noData, exp: ErrNoData{Table: tn}},
		{name: "with data, no users", store: withData, exp: ErrTableNotFound{}},
		{name: "with users", store: withUsers, exp: nil},
	}

	for _, tc := range tcs {

		t.Run(tc.name, func(t *testing.T) {

			_, err := tc.store.All(tn)

			ok := errors.Is(err, tc.exp)

			if !ok {
				t.Fatalf("expected error %T, got %T", tc.exp, err)
			}

		})
	}

}
