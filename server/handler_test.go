package server

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/arkrozycki/reunion/user"
	"github.com/stretchr/testify/require"
)

type test struct {
	name     string
	method   string
	uri      string
	reqBody  interface{}
	respCode int
	respBody *string
}

func CheckExpected(tc test, t *testing.T, w *httptest.ResponseRecorder) {
	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("test: %s, got %v", tc.name, err)
	}
	if res.StatusCode != tc.respCode {
		t.Fatalf("test: %s, status code expected %d got %d", tc.name, tc.respCode, res.StatusCode)
	}
	if tc.respBody != nil {
		if string(body) != *tc.respBody {
			t.Fatalf("test: %s, body expected %s, got %s", tc.name, *tc.respBody, body)
		}
	}
}

func TestHandleRoot(t *testing.T) {
	tests := []test{
		{
			name:     "root",
			method:   http.MethodGet,
			uri:      "/",
			reqBody:  nil,
			respCode: http.StatusOK,
		},
	}

	for _, tc := range tests {
		var body []byte
		if tc.reqBody == nil {
			body = nil
		} else {
			body, _ = json.Marshal(tc.reqBody)
		}
		r := httptest.NewRequest(tc.method, tc.uri, bytes.NewReader(body))
		w := httptest.NewRecorder()
		HandleRoot(w, r)
		CheckExpected(tc, t, w)
	}
}

func TestHandleUserRegistration(t *testing.T) {
	errBadRequest := user.ErrBadRequest.Error()
	errExists := user.ErrExists.Error()
	tests := []test{
		{
			name:     "garbage body",
			method:   http.MethodPost,
			uri:      "/register",
			reqBody:  map[string]interface{}{"hi": "im garbage"},
			respCode: http.StatusBadRequest,
			respBody: &errBadRequest,
		},
		{
			name:     "no user",
			method:   http.MethodPost,
			uri:      "/register",
			reqBody:  user.User{},
			respCode: http.StatusBadRequest,
			respBody: &errBadRequest,
		},
		{
			name:   "email only",
			method: http.MethodPost,
			uri:    "/register",
			reqBody: user.User{
				Email: "test@test.com",
			},
			respCode: http.StatusBadRequest,
			respBody: &errBadRequest,
		},
		{
			name:     "email and password",
			method:   http.MethodPost,
			uri:      "/register",
			reqBody:  user.User{Email: "test@test.com", Password: "mypasspass"},
			respCode: http.StatusOK,
		},
		{
			name:     "duplicate register",
			method:   http.MethodPost,
			uri:      "/register",
			reqBody:  user.User{Email: "test@test.com", Password: "mypasspass"},
			respCode: http.StatusBadRequest,
			respBody: &errExists,
		},
	}

	for _, tc := range tests {
		var body []byte
		if tc.reqBody == nil {
			body = nil
		} else {
			body, _ = json.Marshal(tc.reqBody)
		}
		r := httptest.NewRequest(tc.method, tc.uri, bytes.NewReader(body))
		w := httptest.NewRecorder()
		HandleUserRegistration(w, r)
		CheckExpected(tc, t, w)
	}
}

// TestHandleUser
func TestHandleUser(t *testing.T) {
	tests := []test{
		{
			name:     "get no user",
			method:   http.MethodGet,
			uri:      "/user?id=123",
			reqBody:  nil,
			respCode: http.StatusNotFound,
		},
	}

	for _, tc := range tests {
		var body []byte
		if tc.reqBody == nil {
			body = nil
		} else {
			body, _ = json.Marshal(tc.reqBody)
		}
		r := httptest.NewRequest(tc.method, tc.uri, bytes.NewReader(body))
		w := httptest.NewRecorder()
		HandleUser(w, r)
		CheckExpected(tc, t, w)
	}
}

func TestSend404(t *testing.T) {
	w := httptest.NewRecorder()
	Send404(w)

	res := w.Result()
	if res.StatusCode != http.StatusNotFound {
		t.Fatalf("test: %s, expected %d got %d", "check Send404", 404, res.StatusCode)
	}
}

func TestSendErr(t *testing.T) {
	w := httptest.NewRecorder()
	testErr := errors.New("test error")
	SendErr(w, testErr)
	res := w.Result()
	if res.StatusCode != http.StatusInternalServerError {
		t.Fatalf("test: %s, expected %d got %d", "check senderr", 500, res.StatusCode)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if string(body) != testErr.Error() {
		t.Fatalf("test: %s, expected %s got %s", "sent error", testErr.Error(), string(body))
	}
}

func TestSendJSONResponse(t *testing.T) {
	w := httptest.NewRecorder()
	temp := map[string]interface{}{"test": "sent"}
	SendJSONResponse(w, temp, nil)

	res := w.Result()
	if res.StatusCode != http.StatusOK {
		t.Fatalf("test: %s, expected %d got %d", "send json", 200, res.StatusCode)
	}
	body, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	jtemp, _ := json.Marshal(temp)
	require.JSONEq(t, string(jtemp), string(body))

	testErr := errors.New("sample error")
	w2 := httptest.NewRecorder()
	SendJSONResponse(w2, temp, testErr)
	res2 := w2.Result()
	if res2.StatusCode != http.StatusInternalServerError {
		t.Fatalf("test: %s, expected %d got %d", "check senderr", 500, res.StatusCode)
	}
}
