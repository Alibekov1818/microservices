package data

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"phones_service/internal/validator"
	"time"
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

func (m PhoneModel) Insert(phone *Phone) error {
	query := `INSERT INTO phones (model, brand, year, price)
				VALUES ($1, $2, $3, $4)
				RETURNING id`
	args := []any{phone.Model, phone.Brand, phone.Year, phone.Price}
	return m.DB.QueryRow(query, args...).Scan(&phone.ID)
}

func (m PhoneModel) Get(id int64) (*Phone, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT *
		FROM phones
		WHERE id = $1`

	var phone Phone

	err := m.DB.QueryRow(query, id).Scan(
		&phone.ID,
		&phone.Model,
		&phone.Brand,
		&phone.Year,
		&phone.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &phone, nil
}

func (m PhoneModel) Update(phone *Phone) error {
	query := `
			UPDATE phones
			SET model = $1, brand = $2, year = $3, price = $4
			WHERE id = $5
			RETURNING id`
	args := []any{
		phone.Model,
		phone.Brand,
		phone.Year,
		phone.Price,
		phone.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&phone.ID)
}

func (m PhoneModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
				DELETE FROM phones
				WHERE id = $1`
	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func (m PhoneModel) SelectAll() ([]*Phone, error) {
	query := fmt.Sprintf(`
								SELECT *
								FROM phones ORDER BY id`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	phones := []*Phone{}
	for rows.Next() {
		var phone Phone

		err := rows.Scan(
			&phone.ID,
			&phone.Model,
			&phone.Brand,
			&phone.Year,
			&phone.Price,
		)
		if err != nil {
			return nil, err
		}
		phones = append(phones, &phone)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return phones, nil
}
