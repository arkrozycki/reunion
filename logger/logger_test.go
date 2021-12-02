package logger

import (
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestGet(t *testing.T) {
	l := Get()
	z := zerolog.New(zerolog.ConsoleWriter{})

	if reflect.TypeOf(l) != reflect.TypeOf(z) {
		t.Fail()
	}
}
