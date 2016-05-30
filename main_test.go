package main

import (
	"reflect"
	"testing"
)

func TestVersionString(t *testing.T) {
	typ := reflect.TypeOf(VERSION)
	if typ.String() != "string" {
		t.Errorf("expected VERSION to be a string, got %#v (type %#v)", VERSION, typ.String())
	}
}
