package controllers

import "testing"

func TestSimple(t *testing.T) {
	if 1+1 != 2 {
		t.Error(`nyan.Meowify("cats") != "nyan"`)
	}
}
