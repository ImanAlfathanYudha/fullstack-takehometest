package transport

import (
	"backend/internal/entity"
	"backend/internal/openapigen"
	"encoding/json"
	"fmt"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func ConvertPaymentIntoOpenAPIVersion(p entity.Payment) openapigen.Payment {
	amt := fmt.Sprintf("%.2f", p.Amount)
	return openapigen.Payment{
		Id:        &p.ID,
		Merchant:  &p.Merchant,
		Status:    &p.Status,
		Amount:    &amt,
		CreatedAt: &p.CreatedAt,
	}
}

func ConvertToListPayments(payments []entity.Payment) []openapigen.Payment {
	result := make([]openapigen.Payment, 0, len(payments))
	for _, p := range payments {
		result = append(result, ConvertPaymentIntoOpenAPIVersion(p))
	}
	return result
}
