package handler

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/afthaab/job-portal/internal/middleware"
// 	"github.com/afthaab/job-portal/internal/models"
// 	"github.com/afthaab/job-portal/internal/service"
// 	"github.com/afthaab/job-portal/internal/service/mockmodels"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// 	"go.uber.org/mock/gomock"
// )

// func Test_handler_Signin(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "error in validating the json body",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"name":"",
// 				"email":    "afthab@gmail.com",
// 				"password": afthab}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide valid email and password"}`,
// 		},
// 		{
// 			name: "service layer error",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"name":"",
// 				"email":    "afthab@gmail.com",
// 				"password": "afthab"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().UserSignIn(c.Request.Context(), gomock.Any()).Return("", errors.New("test error from mock function")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"test error from mock function"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"username":"afthab",
// 					"email":    "afthab@gmail.com",
// 					"password": "afthab"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().UserSignIn(c.Request.Context(), gomock.Any()).Return("test token generated", nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"token":"test token generated"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := handler{
// 				service: ms,
// 			}

// 			h.Signin(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// func Test_handler_SignUp(t *testing.T) {
// 	tests := []struct {
// 		name               string
// 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// 		expectedStatusCode int
// 		expectedResponse   string
// 	}{
// 		{
// 			name: "missing trace id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"Internal Server Error"}`,
// 		},
// 		{
// 			name: "error in validating the json body",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"username":"afthab",
// 				"email":    "afthab@gmail.com",
// 				"password": afthab}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"please provide valid username, email and password"}`,
// 		},
// 		{
// 			name: "service layer error",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"username":"afthab",
// 				"email":    "afthab@gmail.com",
// 				"password": "afthab"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().UserSignup(c.Request.Context(), gomock.Any()).Return(models.User{}, errors.New("test error from mock function")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"test error from mock function"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", strings.NewReader(`{"username":"afthab",
// 				"email":    "afthab@gmail.com",
// 				"password": "afthab"}`))
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().UserSignup(c.Request.Context(), gomock.Any()).Return(models.User{
// 					Username:     "afthab",
// 					Email:        "afthab606@gmail.com",
// 					PasswordHash: "hashedpass",
// 				}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"username":"afthab","email":"afthab606@gmail.com"}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := handler{
// 				service: ms,
// 			}

// 			h.SignUp(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }
