package data

import (
	"database/sql"
	"electronicstore/internal/validator"
)

type Phone struct {
	ID    int64  `json:"id"`
	Model string `json:"model"`
	Brand string `json:"brand"`
	Year  int64  `json:"year"`
	Price int64  `json:"price"`
}

func ValidatePhone(v *validator.Validator, phone *Phone) {
	v.Check(phone.Model != "", "model", "must be provided")
	v.Check(phone.Brand != "", "brand", "must be provided")
	v.Check(phone.Year != 0, "year", "must be provided")
	v.Check(phone.Price != 0, "price", "must be provided")
}

type PhoneModel struct {
	DB *sql.DB
}
