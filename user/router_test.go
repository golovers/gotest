package user_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golovers/gotest/user"

	gomock "github.com/golang/mock/gomock"
)

func TestHandleRegisterWithServer(t *testing.T) {
	srv := NewMockservice(gomock.NewController(t))
	handler := user.NewHandler(srv)
	router := user.NewRouter(handler)
	user1 := &user.User{
		ID:   "123",
		Name: "jack",
	}
	serverErr := errors.New("internal server error")
	type expect struct {
		code int
		user user.User
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
				user: user.User{
					ID:   "123",
					Name: "jack",
				},
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
			},
		},
	}
	server := httptest.NewServer(router)
	defer server.Close()
	client := http.Client{
		Timeout: time.Second,
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			r, err := http.NewRequest(http.MethodPost, server.URL+"/users", strings.NewReader(`{"id":"123","name":"jack"}`))
			if err != nil {
				t.Error(err)
			}
			res, err := client.Do(r)
			if err != nil {
				t.Error(err)
			}
			if res.StatusCode != test.expect.code {
				t.Errorf("got code=%d, wants code=%d", res.StatusCode, test.expect.code)
				return
			}
			if res.StatusCode != http.StatusOK {
				return
			}
			var gotUser user.User
			if err := json.NewDecoder(res.Body).Decode(&gotUser); err != nil {
				t.Error(err)
				return
			}
			res.Body.Close()

			if gotUser.ID != test.expect.user.ID {
				t.Errorf("got user_id=%s, wants user_id=%s", gotUser.ID, test.expect.user.ID)
			}
		})
	}
}
