package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestAdminAuthController_CheckAdminAuth(t *testing.T) {
	tests := []struct {
		name           string
		expectedCode   int
		expectedBody   string
	}{
		{
			name:         "管理者権限確認成功",
			expectedCode: http.StatusOK,
			expectedBody: `{"admin":true,"message":"Admin access granted"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/v1/admin/auth/check", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			controller := NewAdminAuthController()
			err := controller.CheckAdminAuth(c)

			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, rec.Code)
			assert.JSONEq(t, tt.expectedBody, rec.Body.String())
		})
	}
}