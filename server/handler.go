package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/arkrozycki/reunion/user"
)

// HandleRoot method
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("HandleRoot()")
	Send200(w)
}

// HandleUserRegistration method
func HandleUserRegistration(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("HandleUserRegistration()")
	var u user.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		SendBadRequest(w, err)
		return
	}
	if _, err = u.Register(); err != nil {
		SendBadRequest(w, err)
		return
	}

	SendJSONResponse(w, u, err)

}

// HandleUser method
func HandleUser(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("HandleUser")
	var u user.User
	q := r.URL.Query()
	id := q.Get("id")
	if id != "" {
		u.ID = id
	}
	found, err := u.Get()
	if err != nil {
		SendErr(w, err)
		return
	}
	if !found {
		Send404(w)
		return
	}
	SendJSONResponse(w, u, nil)

}

// HandleUserVerify method
func HandleUserVerify(w http.ResponseWriter, r *http.Request) {
	log.Debug().Msg("HandleUserVerify")
	var err error
	var u user.User
	q := r.URL.Query()
	id := q.Get("id")
	code := q.Get("code")
	if id == "" || code == "" {
		SendBadRequest(w, errors.New("missing id and/or code"))
		return
	}

	// get the user record and make sure it exists
	u.ID = id
	found, err := u.Get()
	if err != nil {
		SendErr(w, err)
		return
	}
	if !found {
		Send404(w)
		return
	}

	// verify code
	err = u.ValidateCode(code)
	if err != nil {
		SendErr(w, err)
		return
	}
	// code is verifed, update user status
	err = u.SetVerified()
	if err != nil {
		SendErr(w, err)
		return
	}
	Send200(w)

}

// Send404 method
func Send404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

// SendJSONResponse method
func SendJSONResponse(w http.ResponseWriter, obj interface{}, err error) {
	if err != nil {
		SendErr(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(obj)
}

// SendErr method
func SendErr(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err != nil {
		// json.NewEncoder(w).Encode(err)
		fmt.Fprintf(w, err.Error())
	}
}

// SendBadRequest method
func SendBadRequest(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if err != nil {
		// json.NewEncoder(w).Encode(err)
		fmt.Fprintf(w, err.Error())
	}
}

// Send200 method
func Send200(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}
