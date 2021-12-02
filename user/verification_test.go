package user

import (
	"testing"

	"github.com/arkrozycki/reunion/datastore"
)

type testFunc func()

type verifyTest struct {
	name       string
	id         string
	code       string
	err        error
	execBefore testFunc
	execAfter  testFunc
}

func TestValidateCode(t *testing.T) {
	id := "a_test_id"
	code, err := SaveVerificationCode(id)
	if err != nil {
		t.Fatalf("test: validateCode failed with %+v", err)
	}

	tests := []verifyTest{
		{
			name: "no record found",
			id:   "id_not_exist",
			code: code,
			err:  datastore.ErrNotFound,
		},
		{
			name: "code mismatch",
			id:   id,
			code: "not_acutal_code",
			err:  ErrCodeMismatch,
		},
		{
			name: "record found",
			id:   id,
			code: code,
			err:  nil,
		},
		{
			name: "already verified",
			id:   id,
			code: code,
			err:  ErrAlreadyVerified,
		},
		{
			name: "expired code",
			id: "expire_id",
			code: "expired",
			err: ErrCodeExpired,
			execBefore: func() {
				AddExpirationSecs = int64(-1)
				SaveVerificationCode("expire_id")
			},
			execAfter: func() {
			},
		},
	}

	for _, tc := range tests {
		if tc.execBefore != nil {
			tc.execBefore()
		}
		err := validateCode(tc.id, tc.code)
		if err != tc.err {
			t.Fatalf("test: %s expect %s got %s", tc.name, tc.err, err)
		}
	}
}
