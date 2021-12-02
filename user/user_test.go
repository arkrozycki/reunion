package user

import (
	"testing"
)

type test struct {
	name      string
	user      User
	code      string
	err       error
	afterFunc testUser
}

type testUser func(string, User)

func TestRegister(t *testing.T) {
	tests := []test{
		{"all empty", User{}, "", ErrBadRequest, nil},
		{"only email", User{Email: "test@test.com"}, "", ErrBadRequest, nil},
		{"email and pass", User{Email: "test@test.com", Password: "mypassword"}, "", nil, nil},
		{"duplicate", User{Email: "test@test.com", Password: "mypassword"}, "", ErrExists, nil},
	}

	for _, tc := range tests {
		code, err := tc.user.Register()
		if tc.err != err {
			t.Fatalf("test: %s, expected %+v got %+v", tc.name, tc.err, err)
		}
		if tc.afterFunc != nil {
			tc.afterFunc(code, tc.user)
		}
	}
}

func TestUserValidateCode(t *testing.T) {
	user := User{
		ID:          "",
		Email:       "testcodes@test.com",
		FirstName:   "",
		LastName:    "",
		DisplayName: "",
		Password:    "mypassword",
		Role:        0,
		Status:      0,
		ChannelIds:  []int64{},
	}
	code, _ := user.Register()
	tests := []test{
		{
			name: "bad code",
			user: user,
			code: "my_bad",
			err:  ErrCodeMismatch,
		},
		{
			name: "good code",
			user: user,
			code: code,
			err:  nil,
		},
	}

	for _, tc := range tests {
		err := user.ValidateCode(tc.code)
		if tc.err != err {
			t.Fatalf("test: %s, expected %+v got %+v", tc.name, tc.err, err)
		}
		if tc.afterFunc != nil {
			tc.afterFunc(code, tc.user)
		}
	}
}

func TestSetVerified(t *testing.T) {
	user := User{
		ID:          "",
		Email:       "testverified@test.com",
		FirstName:   "",
		LastName:    "",
		DisplayName: "",
		Password:    "mypassword",
		Role:        0,
		Status:      0,
		ChannelIds:  []int64{},
	}
	code, _ := user.Register()
	user.ValidateCode(code)

	err := user.SetVerified()
	if err != nil {
		t.Fatalf("test: SetVerified returned err %+v", err)
	}
	if user.Status != Verified {
		t.Fatalf("test: user status not updated to Verified")
	}
}
