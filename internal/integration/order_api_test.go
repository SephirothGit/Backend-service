package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SephirothGit/Backend-service/internal/handler"
	"github.com/SephirothGit/Backend-service/internal/repository"
	"github.com/SephirothGit/Backend-service/internal/service"
	"github.com/go-chi/chi/v5"
)

func SetupServer() http.Handler {
	repo := repository.NewInMemoryOrderRepository()
	svc := service.NewOrderService(repo)

	r := chi.NewRouter()

	r.Put("/api/v1/orders/{id}", handler.NewOrderHandler(svc))

	return r
}

func TestOrderAPI_UpdateStatus(t *testing.T) {

	tests := []struct {
		name string
		orderID string
		body interface{}
		expectedCode int
	}{
		{
			name:    "invalid json",
			orderID: "123",
			body:    "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid transition",
			orderID: "123",
			body: map[string]string{
				"status": "shipped",
			},
			expectedCode: http.StatusConflict,
		},
		{
			name: "success",
			orderID: "123",
			body: map[string]string{
				"status": "paid",
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name: "order not found",
			orderID: "999",
			body: map[string]string{
				"status": "paid",
			},
			expectedCode: http.StatusNotFound,
		},
	}

	server := SetupServer()

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			var body []byte

			switch v := tt.body.(type) {
			case string:
				body = []byte(v)

			default:
				body, _ = json.Marshal(v)
			}

			req := httptest.NewRequest(
				http.MethodPut,
				"/api/v1/orders/"+tt.orderID,
				bytes.NewBuffer(body),
			)

			rec := httptest.NewRecorder()

			server.ServeHTTP(rec, req)

			if rec.Code != tt.expectedCode {
				t.Fatalf("expected %v got %v", tt.expectedCode, rec.Code)
			}
		})
	}
}