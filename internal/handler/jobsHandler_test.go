package handler

// import (
// 	"context"
// 	"errors"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/afthaab/job-portal/internal/auth"
// 	"github.com/afthaab/job-portal/internal/middleware"
// 	"github.com/afthaab/job-portal/internal/models"
// 	"github.com/afthaab/job-portal/internal/service"
// 	"github.com/afthaab/job-portal/internal/service/mockmodels"
// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang-jwt/jwt/v5"
// 	"go.uber.org/mock/gomock"
// )

// func Test_handler_ViewJobByID(t *testing.T) {
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
// 			name: "missing jwt claims",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest

// 				return c, rr, nil
// 			},
// 			expectedStatusCode: http.StatusUnauthorized,
// 			expectedResponse:   `{"error":"Unauthorized"}`,
// 		},
// 		{
// 			name: "invalid job id",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})

// 				c.Request = httpRequest

// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusBadRequest,
// 			expectedResponse:   `{"error":"Bad Request"}`,
// 		},
// 		{
// 			name: "error while fetching jobs from service",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().ViewJobById(c.Request.Context(), gomock.Any()).Return(models.Jobs{}, errors.New("test service error")).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusInternalServerError,
// 			expectedResponse:   `{"error":"test service error"}`,
// 		},
// 		{
// 			name: "success",
// 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// 				rr := httptest.NewRecorder()
// 				c, _ := gin.CreateTestContext(rr)
// 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// 				ctx := httpRequest.Context()
// 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// 				httpRequest = httpRequest.WithContext(ctx)
// 				c.Request = httpRequest
// 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// 				mc := gomock.NewController(t)
// 				ms := mockmodels.NewMockUserService(mc)

// 				ms.EXPECT().ViewJobById(c.Request.Context(), gomock.Any()).Return(models.Jobs{}, nil).AnyTimes()

// 				return c, rr, ms
// 			},
// 			expectedStatusCode: http.StatusOK,
// 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"cid":0,"name":"","salary":"","notice_period":""}`,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gin.SetMode(gin.TestMode)
// 			c, rr, ms := tt.setup()

// 			h := handler{
// 				service: ms,
// 			}

// 			h.ViewJobByID(c)
// 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// 		})
// 	}
// }

// // func Test_handler_ViewAllJobs(t *testing.T) {
// // 	tests := []struct {
// // 		name               string
// // 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// // 		expectedStatusCode int
// // 		expectedResponse   string
// // 	}{
// // 		{
// // 			name: "missing trace id",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"Internal Server Error"}`,
// // 		},
// // 		{
// // 			name: "missing jwt claims",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusUnauthorized,
// // 			expectedResponse:   `{"error":"Unauthorized"}`,
// // 		},
// // 		{
// // 			name: "error in database",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().ViewAllJobs(c.Request.Context()).Return(nil, errors.New("test error from mock function"))

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"test error from mock function"}`,
// // 		},
// // 		{
// // 			name: "success from database",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().ViewAllJobs(c.Request.Context()).Return([]models.Jobs{
// // 					{
// // 						Cid:          1,
// // 						Name:         "Infosys",
// // 						Salary:       "40000",
// // 						NoticePeriod: "30",
// // 					},
// // 				}, nil)

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusOK,
// // 			expectedResponse:   `[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"cid":1,"name":"Infosys","salary":"40000","notice_period":"30"}]`,
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			gin.SetMode(gin.TestMode)
// // 			c, rr, ms := tt.setup()

// // 			h := handler{
// // 				service: ms,
// // 			}
// // 			h.ViewAllJobs(c)
// // 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// // 			assert.Equal(t, tt.expectedResponse, rr.Body.String())

// // 		})
// // 	}
// // }

// // func Test_handler_ViewJob(t *testing.T) {
// // 	tests := []struct {
// // 		name               string
// // 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// // 		expectedStatusCode int
// // 		expectedResponse   string
// // 	}{
// // 		{
// // 			name: "missing trace id",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"Internal Server Error"}`,
// // 		},
// // 		{
// // 			name: "missing jwt claims",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusUnauthorized,
// // 			expectedResponse:   `{"error":"Unauthorized"}`,
// // 		},
// // 		{
// // 			name: "invalid company id",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusBadRequest,
// // 			expectedResponse:   `{"error":"Bad Request"}`,
// // 		},
// // 		{
// // 			name: "error while fetching job from service",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Request = httpRequest
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().ViewJob(c.Request.Context(), gomock.Any()).Return(nil, errors.New("test service error")).AnyTimes()

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"test service error"}`,
// // 		},
// // 		{
// // 			name: "success case",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Request = httpRequest
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().ViewJob(c.Request.Context(), gomock.Any()).Return([]models.Jobs{
// // 					{
// // 						Cid:          1,
// // 						Name:         "Infosys",
// // 						Salary:       "40000",
// // 						NoticePeriod: "30",
// // 					},
// // 				}, nil).AnyTimes()

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusOK,
// // 			expectedResponse:   `[{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"cid":1,"name":"Infosys","salary":"40000","notice_period":"30"}]`,
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			gin.SetMode(gin.TestMode)
// // 			c, rr, ms := tt.setup()

// // 			h := handler{
// // 				service: ms,
// // 			}
// // 			h.ViewJob(c)
// // 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// // 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// // 		})
// // 	}
// // }

// // func Test_handler_AddJobs(t *testing.T) {
// // 	tests := []struct {
// // 		name               string
// // 		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.UserService)
// // 		expectedStatusCode int
// // 		expectedResponse   string
// // 	}{
// // 		{
// // 			name: "missing trace id",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"Internal Server Error"}`,
// // 		},
// // 		{
// // 			name: "missing jwt claims",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Request = httpRequest

// // 				return c, rr, nil
// // 			},
// // 			expectedStatusCode: http.StatusUnauthorized,
// // 			expectedResponse:   `{"error":"Unauthorized"}`,
// // 		},
// // 		{
// // 			name: "invalid job id",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", nil)
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "abc"})

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusBadRequest,
// // 			expectedResponse:   `{"error":"Bad Request"}`,
// // 		},
// // 		{
// // 			name: "error in validating the json",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"cid":"1",
// // 				"name":    "junior web developer",
// // 				"field": "it}`))
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})
// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusBadRequest,
// // 			expectedResponse:   `{"error":"please provide valid name, location and field"}`,
// // 		},
// // 		{
// // 			name: "error in service layer",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"cid":1,
// // 				"name": "junior web developer",
// // 				"field": "it"}`))
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.Jobs{}, errors.New("test error from mock function"))

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusInternalServerError,
// // 			expectedResponse:   `{"error":"test error from mock function"}`,
// // 		},
// // 		{
// // 			name: "success case",
// // 			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.UserService) {
// // 				rr := httptest.NewRecorder()
// // 				c, _ := gin.CreateTestContext(rr)
// // 				httpRequest, _ := http.NewRequest(http.MethodGet, "http://test.com:8080", strings.NewReader(`{"cid":1,
// // 				"name": "junior web developer",
// // 				"field": "it"}`))
// // 				ctx := httpRequest.Context()
// // 				ctx = context.WithValue(ctx, middleware.TraceIDKey, "123")
// // 				ctx = context.WithValue(ctx, auth.Key, jwt.RegisteredClaims{})
// // 				httpRequest = httpRequest.WithContext(ctx)
// // 				c.Params = append(c.Params, gin.Param{Key: "id", Value: "123"})

// // 				c.Request = httpRequest

// // 				mc := gomock.NewController(t)
// // 				ms := mockmodels.NewMockUserService(mc)

// // 				ms.EXPECT().AddJobDetails(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.Jobs{
// // 					Cid:          1,
// // 					Name:         "Infosys",
// // 					Salary:       "10000",
// // 					NoticePeriod: "30",
// // 				}, nil)

// // 				return c, rr, ms
// // 			},
// // 			expectedStatusCode: http.StatusOK,
// // 			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"cid":1,"name":"Infosys","salary":"10000","notice_period":"30"}`,
// // 		},
// // 	}
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			gin.SetMode(gin.TestMode)
// // 			c, rr, ms := tt.setup()

// // 			h := handler{
// // 				service: ms,
// // 			}
// // 			h.AddJobs(c)
// // 			assert.Equal(t, tt.expectedStatusCode, rr.Code)
// // 			assert.Equal(t, tt.expectedResponse, rr.Body.String())
// // 		})
// // 	}
// // }
