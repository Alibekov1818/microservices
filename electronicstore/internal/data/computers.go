package data

import (
	"context"
	"database/sql"
	"electronicstore/internal/validator"
	"errors"
	"fmt"
	"time"
)

type Computer struct {
	ID     int64  `json:"id"`
	Model  string `json:"model"`
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
	Price  int64  `json:"price"`
}

func ValidateComputer(v *validator.Validator, computer *Computer) {
	v.Check(computer.Model != "", "model", "must be provided")
	v.Check(computer.Cpu != "", "cpu", "must be provided")
	v.Check(computer.Memory != "", "memory", "must be provided")
	v.Check(computer.Price != 0, "price", "must be provided")
}

type ComputerModel struct {
	DB *sql.DB
}

func (m ComputerModel) Insert(computer *Computer) error {
	query := `INSERT INTO computers (model, cpu, memory, price)
				VALUES ($1, $2, $3, $4)
				RETURNING id`
	args := []any{computer.Model, computer.Cpu, computer.Memory, computer.Price}
	return m.DB.QueryRow(query, args...).Scan(&computer.ID)
}

func (m ComputerModel) Get(id int64) (*Computer, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}
	query := `
		SELECT *
		FROM computers
		WHERE id = $1`
	var computer Computer

	err := m.DB.QueryRow(query, id).Scan(
		&computer.ID,
		&computer.Model,
		&computer.Cpu,
		&computer.Memory,
		&computer.Price,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &computer, nil
}

func (m ComputerModel) Update(computer *Computer) error {
	query := `
			UPDATE computers
			SET model = $1, cpu = $2, memory = $3, price = $4
			WHERE id = $5
			RETURNING id`
	args := []any{
		computer.Model,
		computer.Cpu,
		computer.Memory,
		computer.Price,
		computer.ID,
	}
	return m.DB.QueryRow(query, args...).Scan(&computer.ID)
}

func (m ComputerModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `
				DELETE FROM computers
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

func (m ComputerModel) SelectAll() ([]*Computer, error) {
	query := fmt.Sprintf(`
								SELECT *
								FROM computers ORDER BY id`)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	computers := []*Computer{}
	for rows.Next() {
		var computer Computer

		err := rows.Scan(
			&computer.ID,
			&computer.Model,
			&computer.Cpu,
			&computer.Memory,
			&computer.Price,
		)
		if err != nil {
			return nil, err
		}
		computers = append(computers, &computer)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return computers, nil
}
