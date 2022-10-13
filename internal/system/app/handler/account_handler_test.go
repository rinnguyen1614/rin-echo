package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v4"
	"github.com/labstack/echo/v4"
	echox "github.com/rinnguyen1614/rin-echo/internal/core/echo"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/model/request"
	"github.com/rinnguyen1614/rin-echo/internal/system/app/service/mocks"
	_ "github.com/rinnguyen1614/rin-echo/internal/system/domain"
	"github.com/rinnguyen1614/rin-echo/internal/system/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAccount_Login(t *testing.T) {
	requestData := request.Login{}
	err := faker.FakeData(&requestData)
	assert.NoError(t, err)

	testCases := []struct {
		name       string
		body       request.Login
		buildStubs func(c echox.Context, service *mocks.AccountService, userService *mocks.UserService)
		expErrCode int // 0 for Success
	}{
		{
			name: "OK",
			body: requestData,
			buildStubs: func(c echox.Context, mockService *mocks.AccountService, userService *mocks.UserService) {
				mockService.
					On("WithContext", c).Return(mockService).
					On("Login", mock.AnythingOfType("request.Login")).
					Return(mock.AnythingOfType("*jwt.Token"), nil)
			},
		},
		{
			name: errors.ErrUserNamePasswordNotMatch.Message(),
			body: requestData,
			buildStubs: func(c echox.Context, mockService *mocks.AccountService, userService *mocks.UserService) {
				mockService.
					On("WithContext", c).Return(mockService).
					On("Login", mock.AnythingOfType("request.Login")).
					Return(nil, errors.ErrUserNamePasswordNotMatch)
			},
			expErrCode: http.StatusBadRequest,
		},
		{
			name: "Invalid",
			body: request.Login{
				Username: "ad",
				Password: "12345",
			},
			expErrCode: http.StatusBadRequest,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mockService := new(mocks.AccountService)
			mockUserService := new(mocks.UserService)
			handler := AccountHandler{
				service:     mockService,
				userService: mockUserService,
			}

			e := newEchoTest()
			data, err := json.Marshal(testCase.body)
			require.NoError(t, err)
			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(data))
			res := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := echox.NewContextx(e.NewContext(req, res))

			if testCase.buildStubs != nil {
				testCase.buildStubs(c, mockService, mockUserService)
			}

			err = handler.Login(c)

			if testCase.expErrCode != 0 {
				require.Error(t, err)
				e.HTTPErrorHandler(err, c)
				assert.Equal(t, testCase.expErrCode, res.Code)
			} else {
				require.NoError(t, err)
			}

			mockService.AssertExpectations(t)
		})
	}
}
