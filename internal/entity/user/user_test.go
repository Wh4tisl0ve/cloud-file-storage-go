package user

import (
	"errors"
	"testing"
	"time"
	"uuid"
)

func TestCreateUser(t *testing.T) {
	tests := map[string]struct {
		username      string
		password      string
		expectedUser  *User
		expectedError *error
	}{
		"успешное создание пользователя": {
			username: "user11",
			password: "user11",
			expectedUser: &User{
				uuid.New(),
				Username{"user11"},
				"user11",
				time.Now(),
			},
		},
		// todo
		/*
			слишком короткое имя, слишком длинное имя
		*/
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			u, err := New(testCase.username, testCase.password)

			if testCase.expectedUser.username.value != u.username.value {
				t.Fatalf("got %q, want %q", u.username.value, testCase.expectedUser.username.value)
			}

			if testCase.expectedError != nil && errors.Is(err, *testCase.expectedError) {
				t.Fatal("missed error type")
			}
		})
	}
}
