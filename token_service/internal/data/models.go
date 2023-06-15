package data

import (
	"database/sql"
)

type Models struct {
	Tokens TokenModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tokens: TokenModel{DB: db},
	}
}
