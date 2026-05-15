package handler

import (
	"backend/internal/entity"
	"backend/internal/module/payment/usecase"
	"backend/internal/openapigen"
	"backend/internal/transport"
	"net/http"
)

type PaymentHandler struct {
	paymentUC usecase.IPaymentUsecase
}

func NewPaymentHandler(paymentUC usecase.IPaymentUsecase) *PaymentHandler {
	return &PaymentHandler{paymentUC: paymentUC}
}

func (h *PaymentHandler) GetDashboardV1Payments(
	w http.ResponseWriter,
	r *http.Request,
	params openapigen.GetDashboardV1PaymentsParams,
) {
	filter := entity.PaymentFilter{}
	if params.Status != nil {
		filter.Status = *params.Status
	}
	if params.Id != nil {
		filter.ID = *params.Id
	}
	payments, err := h.paymentUC.GetListPayments(filter)
	if err != nil {
		transport.WriteError(w, err)
		return
	}

	result := transport.ConvertToListPayments(payments)
	transport.WriteJSON(w, http.StatusOK, openapigen.PaymentListResponse{
		Payments: &result,
	})
}
