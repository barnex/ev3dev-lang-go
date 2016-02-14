package ev3

import (
	"testing"
)

func init() {
	SysFS = "testdata/"
}

func TestOpenMotor(t *testing.T) {
	m, err := OpenMotor("outC")
	if err != nil {
		t.Fatal(err)
	}
	if m.Address() != "outC" {
		t.Errorf("Address: got %q, want %q:", m.Address(), "outC")
	}
}
