package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/popeskul/car-outliers-detection/internal/domain"
	"github.com/popeskul/car-outliers-detection/internal/services/mocks"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func TestCheckAgesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockIService(ctrl)
	mockOutlierService := mocks.NewMockIOutlierService(ctrl)
	mockService.EXPECT().OutlierService().Return(mockOutlierService).AnyTimes()

	testCases := []struct {
		name             string
		inputMachines    []domain.Machine
		mockSetup        func()
		expectedStatus   int
		expectedResponse interface{}
	}{
		{
			name:          "Valid Request",
			inputMachines: []domain.Machine{{ID: "1", Age: "10 years"}},
			mockSetup: func() {
				mockOutlierService.EXPECT().
					DetectOutliers(gomock.Any()).
					Return([]*domain.Machine{{ID: "1", Age: "10 years"}}, nil)
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: []domain.Machine{{ID: "1", Age: "10 years"}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Setup
			tc.mockSetup()
			router := gin.Default()
			h := NewHandler(mockService)
			h.Init(router)

			// Create request
			body, _ := json.Marshal(CheckAgesRequest(tc.inputMachines))
			req, _ := http.NewRequest(http.MethodPost, "/api/check-ages", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Perform request
			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatus {
				t.Errorf("Expected status %v, got %v", tc.expectedStatus, w.Code)
			}

			if tc.expectedResponse != nil {
				var respBody []domain.Machine
				err := json.Unmarshal(w.Body.Bytes(), &respBody)
				if err != nil {
					t.Errorf("Error unmarshalling response body: %v", err)
				}

				expectedResp, _ := tc.expectedResponse.([]domain.Machine)
				if !compareMachineSlices(respBody, expectedResp) {
					t.Errorf("Expected response body %v, got %v", expectedResp, respBody)
				}
			}
		})
	}
}

// compareMachineSlices compares two slices of domain.Machine for equality.
func compareMachineSlices(a, b []domain.Machine) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].ID != b[i].ID || a[i].Age != b[i].Age {
			return false
		}
	}
	return true
}
