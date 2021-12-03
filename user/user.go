// Package user for user related stuff
package user

import (
	"encoding/json"
	"errors"

	"github.com/arkrozycki/reunion/datastore"
	"github.com/arkrozycki/reunion/logger"
	"github.com/arkrozycki/reunion/util"
)

// some vars of course
var (
	db            = datastore.GetDatastore()
	log           = logger.Get()
	ErrExists     = errors.New("record exists")
	ErrBadRequest = errors.New("bad request")
)

// UserType int32 enum
type UserType int32

// UserStatus int32 enum
type UserStatus int32

// const UserType
const (
	Public UserType = 0
	Admin  UserType = 1
)

// const UserStatus
const (
	Unverified UserStatus = 0
	Verified   UserStatus = 1
	Disabled   UserStatus = 2
)

// User struct
type User struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	DisplayName string     `json:"display_name,omitempty"`
	Password    string     `json:"password,omitempty"`
	Role        UserType   `json:"role,omitempty"`
	Status      UserStatus `json:"status"`
	ChannelIds  []int64    `json:"channel_ids,omitempty"`
}

// Register method
func (u *User) Register() (string, error) {
	log.Debug().Msgf("Register() %+v", u)
	if u.Email == "" {
		return "", ErrBadRequest
	}
	if u.Password == "" {
		return "", ErrBadRequest
	}

	// generate the email hash key
	u.ID = util.Hash(u.Email)

	// check if user key already exists
	if exists, err := u.Get(); exists || err != nil {
		if exists {
			return "", ErrExists
		}
		return "", err
	}

	// additional defaults
	u.Role = Public
	u.Status = Unverified

	// persist to store
	if err := db.SaveUser(u.ID, u); err != nil {
		return "", err
	}

	// generate verfication code
	code, err := SaveVerificationCode(u.ID)
	if err != nil {
		return "", err
	}
	return code, nil

}

// Get method
// retrieve the user by id
func (u *User) Get() (bool, error) {
	log.Debug().Msgf("user.Get() %s", u.ID)

	rec, err := db.GetUser(u.ID)
	if err != nil {
		return false, err
	}
	if rec == "" {
		return false, nil
	}
	err = json.Unmarshal([]byte(rec), &u)
	if err != nil {
		return false, err
	}
	return true, nil

}

// ValidateCode method
// validates a users code
func (u *User) ValidateCode(code string) error {
	var err error
	log.Debug().Msgf("user.ValidateCode() %s", code)

	err = validateCode(u.ID, code)
	if err != nil {
		return err
	}

	return nil
}

// SetVerified method
func (u *User) SetVerified() error {
	log.Debug().Msgf("user.SetVerified() %s", u.ID)
	u.Status = Verified
	if err := db.SaveUser(u.ID, u); err != nil {
		return err
	}
	return nil
}
