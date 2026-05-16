package usecase

import (
	"errors"
	"testing"

	"backend/internal/entity"
)

// ─── MOCK REPOSITORY ─────────────────────────────────────────
// This simulates the database so we can test the Usecase logic
type MockPaymentRepo struct {
	GetPaymentsFunc func(filter entity.PaymentFilter) ([]entity.Payment, error)
}

func (m *MockPaymentRepo) GetPayments(f entity.PaymentFilter) ([]entity.Payment, error) {
	return m.GetPaymentsFunc(f)
}

func (m *MockPaymentRepo) GetPaymentByID(id string) (*entity.Payment, error) { return nil, nil }

func TestListPayments_Success(t *testing.T) {
	mockData := []entity.Payment{
		{ID: "1", Merchant: "Tokopedia", Status: "completed", Amount: 100},
	}

	mockRepo := &MockPaymentRepo{
		GetPaymentsFunc: func(filter entity.PaymentFilter) ([]entity.Payment, error) {
			return mockData, nil
		},
	}

	// 2. Initialize Usecase with Mock
	uc := NewPaymentUsecase(mockRepo)

	// 3. Execute
	results, err := uc.GetListPayments(entity.PaymentFilter{})

	// 4. Assert
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}
	if results[0].Merchant != "Tokopedia" {
		t.Errorf("Expected Tokopedia, got %s", results[0].Merchant)
	}
}

func TestListPayments_Error(t *testing.T) {
	// 1. Setup mock to return an error
	mockRepo := &MockPaymentRepo{
		GetPaymentsFunc: func(filter entity.PaymentFilter) ([]entity.Payment, error) {
			return nil, errors.New("db connection failed")
		},
	}

	uc := NewPaymentUsecase(mockRepo)
	_, err := uc.GetListPayments(entity.PaymentFilter{})

	if err == nil {
		t.Error("Expected an error from usecase, got nil")
	}
}
