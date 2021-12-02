package user

import (
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

// Verification struct
type Verification struct {
	UserID    string `json:"user_id"`
	Code      string `json:"code"`
	CreatedAt int64  `json:"created_at"`
	ExpiresAt int64  `json:"expires_at"`
	Completed bool
}

const codeLen = 8

var (
	ErrAlreadyVerified = errors.New("code already verified")
	ErrCodeExpired     = errors.New("code expired")
	ErrCodeMismatch    = errors.New("code not matched")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var AddExpirationSecs = int64(86400)

// SaveVerificationCode method
// create a record for the key
func SaveVerificationCode(id string) (string, error) {
	log.Debug().Msgf("SaveVerificationCode %s", id)
	var err error
	code := generateCode(codeLen)
	ts := time.Now().Unix()
	v := Verification{
		UserID:    id,
		Code:      code,
		CreatedAt: ts,
		ExpiresAt: ts + AddExpirationSecs,
	}

	log.Debug().Msgf("save code %s", code)
	err = db.SaveCode(id, v)
	if err != nil {
		return "", err
	}

	return code, nil

}

// validateCode method
func validateCode(id string, code string) error {
	log.Debug().Msgf("validateCode %s for %s", code, id)
	rec, err := db.GetCode(id)
	if err != nil {
		return err
	}

	var v Verification
	err = json.Unmarshal([]byte(rec), &v)
	if err != nil {
		return err
	}

	log.Debug().Msgf("v record: %+v", v)

	if v.Completed {
		return ErrAlreadyVerified
	}

	ts := time.Now().Unix()
	if ts >= v.ExpiresAt {
		return ErrCodeExpired
	}

	if v.Code != code {
		return ErrCodeMismatch
	}

	// if we made it this far, we are good
	err = v.Complete()
	if err != nil {
		return err
	}

	return nil
}

// Complete method
// marks the verification complete
func (v *Verification) Complete() error {
	v.Completed = true
	err := db.SaveCode(v.UserID, v)
	if err != nil {
		return err
	}
	return nil
}

// generateCode method
func generateCode(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
