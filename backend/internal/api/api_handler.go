package api

import (
	"net/http"

	authHandler "backend/internal/module/auth/handler"
	paymentHandler "backend/internal/module/payment/handler"
	"backend/internal/openapigen"
)

type APIHandler struct {
	Auth    *authHandler.AuthHandler
	Payment *paymentHandler.PaymentHandler
}

var _ openapigen.ServerInterface = (*APIHandler)(nil)

func (h *APIHandler) PostDashboardV1AuthLogin(w http.ResponseWriter, r *http.Request) {
	h.Auth.PostDashboardV1AuthLogin(w, r)
}

func (h *APIHandler) GetDashboardV1Payments(w http.ResponseWriter, r *http.Request, params openapigen.GetDashboardV1PaymentsParams) {
	h.Payment.GetDashboardV1Payments(w, r, params)
}
