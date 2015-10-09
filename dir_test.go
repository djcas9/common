package common

import "testing"

func TestIsDir(t *testing.T) {

	if IsDir("file.go") {
		t.Error("file.go is not a directory")
	}

	if !IsDir("/") {
		t.Error("/ should be a directory")
	}

}
