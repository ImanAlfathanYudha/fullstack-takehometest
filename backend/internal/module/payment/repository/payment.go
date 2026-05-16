package repository

import (
	"backend/internal/entity"
	"database/sql"
)

type PaymentRepository struct {
	db *sql.DB
}

func NewPaymentRepo(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

// Interface
type IPaymentRepository interface {
	GetPayments(filter entity.PaymentFilter) ([]entity.Payment, error)
	GetPaymentByID(id string) (*entity.Payment, error)
}

func (r *PaymentRepository) GetPayments(filter entity.PaymentFilter) ([]entity.Payment, error) {
	query := "SELECT id, merchant, status, amount, created_at FROM payments"
	args := []any{}

	if filter.Status != "" {
		query += " WHERE status = ?"
		args = append(args, filter.Status)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, entity.WrapError(
			err,
			entity.ErrorCodeInternal,
			"failed to query payments",
		)
	}
	defer rows.Close()
	var payments []entity.Payment
	for rows.Next() {
		var p entity.Payment
		if err := rows.Scan(&p.ID, &p.Merchant, &p.Status, &p.Amount, &p.CreatedAt); err != nil {
			return nil, entity.WrapError(err, entity.ErrorCodeInternal, "failed to scan payment")
		}
		payments = append(payments, p)
	}
	return payments, nil
}

func (r *PaymentRepository) GetPaymentByID(id string) (*entity.Payment, error) {
	row := r.db.QueryRow(
		"SELECT id, merchant, status, amount, created_at FROM payments WHERE id = ?", id,
	)
	var p entity.Payment
	if err := row.Scan(&p.ID, &p.Merchant, &p.Status, &p.Amount, &p.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, entity.ErrorNotFound("payment not found")
		}
		return nil, entity.WrapError(err, entity.ErrorCodeInternal, "db error")
	}
	return &p, nil
}
