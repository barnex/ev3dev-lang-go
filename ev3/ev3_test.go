package ev3

import (
	"testing"
)

func init() {
	SysFS = "testdata/sys/class/"
}

func TestOpenMotor(t *testing.T) {
	m, err := OpenMotor("outA")
	if err != nil {
		t.Error(err)
	}
	if m.Address() != "outA" {
		t.Errorf("Address: got %v, want %v:", m.Address(), "outA")
	}
}
