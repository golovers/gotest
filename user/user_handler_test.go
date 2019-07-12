package user_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golovers/gotest/user"

	gomock "github.com/golang/mock/gomock"
)

func TestHandleRegister(t *testing.T) {
	srv := NewMockservice(gomock.NewController(t))
	handler := user.NewHandler(srv)
	user1 := &user.User{
		ID:   "123",
		Name: "jack",
	}
	serverErr := errors.New("internal server error")
	type expect struct {
		code int
		body string
	}
	testCases := []struct {
		name     string
		tearDown func()
		input    *user.User
		expect   expect
	}{
		{
			name: "register success",
			tearDown: func() {
				srv.EXPECT().Register(gomock.Any(), user1).Times(1).Return(nil)
			},
			input: user1,
			expect: expect{
				code: http.StatusOK,
				body: `{"id":"123","name":"jack"}`,
			},
		},
		{
			name: "register failed",
			tearDown: func() {
				srv.EXPECT().Register(gomock.Any(), user1).Times(1).Return(serverErr)
			},
			input: user1,
			expect: expect{
				code: http.StatusInternalServerError,
				body: serverErr.Error(),
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodPost, "", strings.NewReader(`{"id":"123","name":"jack"}`))
			if err != nil {
				t.Error(err)
			}
			handler.Register(w, r)
			if w.Code != test.expect.code {
				t.Errorf("got code=%d, wants code=%d", w.Code, test.expect.code)
			}
			gotBody := strings.TrimSpace(w.Body.String())
			if gotBody != test.expect.body {
				t.Errorf("got body=%s, wants body=%s", gotBody, test.expect.body)
			}
		})
	}
}
