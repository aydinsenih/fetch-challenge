package database

import (
	"errors"

	"github.com/google/uuid"
)

var db = make(map[uuid.UUID]int)

func Write(id uuid.UUID, d int) error {
	_, ok := db[id]
	if !ok {
		db[id] = d
		return nil
	}
	return errors.New("Error Writing DB")
}
func Read(id uuid.UUID) (int, bool) {
	i, ok := db[id]
	if !ok {
		return 0, false
	}
	return i, true
}
