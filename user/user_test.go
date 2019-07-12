package user_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golovers/gotest/user"

	gomock "github.com/golang/mock/gomock"
)

func TestRegisterUser(t *testing.T) {
	mockedRepo := NewMockrepository(gomock.NewController(t))
	service := user.NewService(mockedRepo)
	user1 := &user.User{
		ID:   "123",
		Name: "jack",
	}
	dbErr := errors.New("cannot connect to database")
	testCases := []struct {
		name     string
		tearDown func()
		input    *user.User
		ouput    error
	}{
		{
			name:  "user not exist, create succeed",
			input: user1,
			tearDown: func() {
				mockedRepo.EXPECT().Get(gomock.Any(), user1.ID).Times(1).Return(nil, user.ErrNotFound)
				mockedRepo.EXPECT().Create(gomock.Any(), user1).Times(1).Return(nil)
			},
			ouput: nil,
		},
		{
			name:  "user already exist, update success",
			input: user1,
			tearDown: func() {
				mockedRepo.EXPECT().Get(gomock.Any(), user1.ID).Times(1).Return(user1, nil)
				mockedRepo.EXPECT().Update(gomock.Any(), user1).Times(1).Return(nil)
			},
			ouput: nil,
		},
		{
			name:  "user already exist, update failed because of database connection failed",
			input: user1,
			tearDown: func() {
				mockedRepo.EXPECT().Get(gomock.Any(), user1.ID).Times(1).Return(user1, nil)
				mockedRepo.EXPECT().Update(gomock.Any(), user1).Times(1).Return(dbErr)
			},
			ouput: dbErr,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			test.tearDown()
			err := service.Register(context.TODO(), test.input)
			if err != test.ouput {
				t.Errorf("got err = %v, expects err = %v", err, test.ouput)
			}
		})
	}
}
