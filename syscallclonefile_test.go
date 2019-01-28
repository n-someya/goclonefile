package goclonefile

import (
	"os"
	"testing"
)

const (
	dummy_filename  = "./test/dummy.txt"
	copied_filename = "./test/dummy_2.txt"
)

func TestClonefile(t *testing.T) {
	err := Clonefile(dummy_filename, copied_filename)
	if err != nil {
		t.Fail()
	}
	_, err = os.Stat(copied_filename)
	if err != nil {
		t.Fail()
	}
	os.Remove(copied_filename)
}
