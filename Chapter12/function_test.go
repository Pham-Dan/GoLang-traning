package main

import "testing"

func TestMatchInt(t *testing.T) {
	if matchInt("") {
		t.Error(`matchInt("") != false`)
	}

	if matchInt("00") == false {
		t.Error(`matchInt("00") != true`)
	}

	if matchInt("-00") == false {
		t.Error(`matchInt("-00") != true`)
	}

	if matchInt("+00") == false {
		t.Error(`matchInt("+00") != true`)
	}
}
