package server

import (
	"reflect"
	"testing"

	"github.com/gorilla/mux"
)

func TestNew(t *testing.T) {
	r := Router{}
	r.New()
	if r.router == nil {
		t.Fatalf("test: router new failed")
	}

	temp := mux.NewRouter()
	if reflect.TypeOf(r.router) != reflect.TypeOf(temp) {
		t.Fatalf("test: router returned is not mux.Router")
	}

}

func TestAddRoutes(t *testing.T) {
	r := Router{}
	r.New()
	r.AddRoutes()
	//TODO: test something i guess
}
