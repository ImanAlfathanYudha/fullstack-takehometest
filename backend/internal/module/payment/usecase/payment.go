package usecase

import (
	"backend/internal/entity"
	"backend/internal/module/payment/repository"
)

type PaymentUsecase struct {
	repo repository.IPaymentRepository
}

type IPaymentUsecase interface {
	GetListPayments(filter entity.PaymentFilter) ([]entity.Payment, error)
}

func NewPaymentUsecase(repo repository.IPaymentRepository) *PaymentUsecase {
	return &PaymentUsecase{repo: repo}
}
func (u *PaymentUsecase) GetListPayments(filter entity.PaymentFilter) ([]entity.Payment, error) {
	status := filter.Status
	if status != "" && !isValidStatus(status) {
		return nil, entity.ErrorBadRequest("invalid status: must be completed, processing, or failed")
	}
	listPayments, err := u.repo.GetPayments(filter)
	if err != nil {
		return nil, err
	}
	return listPayments, nil
}

func isValidStatus(status string) bool {
	switch status {
	case "completed", "processing", "failed":
		return true
	}
	return false
}
