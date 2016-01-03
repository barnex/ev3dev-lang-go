package ev3

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"sort"
	"strings"
)

// Path to driver files (overridden for tests).
var SysFS = "/sys/class/"

type IODevice struct {
	path string
}

func OpenIODevice(sysClass, devNameConvention, port string) (IODevice, error) {
	sysDir := path.Join(SysFS, sysClass)
	dir, err := os.Open(sysDir)
	if err != nil {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, err)

	}
	ls, err := dir.Readdirnames(-1)
	if err != nil {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, err)
	}

	if len(ls) == 0 {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, "no devices present")
	}

	sort.Strings(ls)
	if port == "" {
		return IODevice{path: path.Join(sysDir, ls[0])}, nil
	}

	for _, f := range ls {
		if readFile(path.Join(sysDir, f, "port_name")) == port {
			return IODevice{path: path.Join(sysDir, f)}, nil
		}
	}
	return IODevice{}, fmt.Errorf("open %v:%v: port not found in %v", sysClass, port, ls)
}

func readFile(f string) string {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%s: %q", f, bytes)
	return strings.TrimSpace(string(bytes))
}

func (d IODevice) writeString(file string, x string) {
	panic("not implemented")
}

func (d IODevice) writeStringSelector(file string, x string) {
	panic("not implemented")
}

func (d IODevice) writeInt(file string, x int) {
	panic("not implemented")
}

func (d IODevice) readInt(file string) int {
	panic("not implemented")
}

func (d IODevice) readString(file string) string {
	panic("not implemented")
}

func (d IODevice) readStringArray(file string) []string {
	panic("not implemented")
}

func (d IODevice) readStringSelector(file string) string {
	panic("not implemented")
}
