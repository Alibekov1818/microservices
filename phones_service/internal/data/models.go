package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Phones PhoneModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Phones: PhoneModel{DB: db},
	}
}
