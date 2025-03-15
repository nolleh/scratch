package store

import "time"

type Model struct {
  tn string
  mds Models
}

type Models []Model

type data map[string]Models

type Store struct {
  data map[string]Models
}

type ErrNoData struct {
	error
	Table string
}

type ErrTableNotFound struct {
	error
	Table      string
	OccurredAt time.Time
}

// Is method checks if the target is of the same type
func (e *ErrTableNotFound) Is(target error) bool {
	_, ok := target.(*ErrTableNotFound)
	return ok
}
func (e *ErrNoData) Is(target error) bool {
	_, ok := target.(*ErrNoData)
	return ok
}

func (s *Store) All(tn string) (Models, error) {
	db := s.data

	if db == nil {
    return nil, ErrNoData{Table:tn}
	}

	mods, ok := db[tn]

	if !ok {
		return nil, ErrTableNotFound{
			Table:      tn,
			OccurredAt: time.Now(),
		}

	}

	return mods, nil
}
