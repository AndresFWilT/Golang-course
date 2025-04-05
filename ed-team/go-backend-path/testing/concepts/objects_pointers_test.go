package main

import (
	"reflect"
	"testing"
)

type Kind struct {
	Name string
}

type Dog struct {
	Name string
	Age  uint8
	Kind Kind
}

func TestObjectAndPointers(t *testing.T) {
	want := Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Criollo",
		},
	}

	got := Dog{
		Name: "Firulais",
		Age:  1,
		Kind: Kind{
			Name: "Criollo",
		},
	}

	// comparing objects
	if want != got {
		t.Errorf("Wanted %v, got %v", want, got)
	}

	// comparing pointers
	wantPointer := &want
	gotPointer := &got

	// Use reflect package to compare pointers and similarities between pointers. Very useful
	if !reflect.DeepEqual(wantPointer, gotPointer) {
		t.Errorf("Wanted %v, got %v", wantPointer, gotPointer)
	}
}
