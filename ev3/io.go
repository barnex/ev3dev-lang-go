package ev3

import (
	"bytes"
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

// sysClass: e.g.: "tacho-motor"
// devNameConvetion: e.g.: "motor{0} (unused)
// port: e.g.: "outC"
func OpenIODevice(sysClass, devNameConvention, port string) (IODevice, error) {

	// list all files under /sys/class/<sysClass>
	sysDir := path.Join(SysFS, sysClass)
	dir, err := os.Open(sysDir)
	if err != nil {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, err)

	}
	ls, err := dir.Readdirnames(-1)
	if err != nil {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, err)
	}
	sort.Strings(ls)
	if len(ls) == 0 {
		return IODevice{}, fmt.Errorf("open %v:%v: %v", sysClass, port, "no devices present")
	}

	// no port specified: return first matching
	if port == "" {
		return IODevice{path: path.Join(sysDir, ls[0])}, nil
	}

	// search port_name files for matching port
	for _, f := range ls {
		if string(readFile(path.Join(sysDir, f, "address"))) == port {
			return IODevice{path: path.Join(sysDir, f)}, nil
		}
	}
	return IODevice{}, fmt.Errorf("open %v:%v: port not found in %v", sysClass, port, ls)
}

func (d *IODevice) String() string {
	return d.path
}

func readFile(f string) []byte {
	bytes, err := ioutil.ReadFile(f)
	if err != nil {
		log.Println(err)
	}
	return clean(bytes)
}

func clean(x []byte) []byte {
	// strip trailing nul characters (test files)
	// TODO(barnex): clean up test files
	for i, b := range x {
		if b == 0 {
			x = x[:i]
			break
		}
	}
	return bytes.Trim(x, "\n")
}

func (d IODevice) write(file string, x interface{}) {
	f, err := os.OpenFile(path.Join(d.path, file), os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(f, x)
}

func (d IODevice) read(file string) []byte {
	f := path.Join(d.path, file)
	return readFile(f)
}

func (d IODevice) writeString(file string, x string) {
	d.write(file, x)
}

func (d IODevice) writeStringSelector(file string, x string) {
	d.write(file, x)
}

func (d IODevice) writeInt(file string, x int) {
	d.write(file, x)
}

func (d IODevice) readInt(file string) int {
	panic("not implemented")
}

func (d IODevice) readString(file string) string {
	return strings.Trim(string(d.read(file)), "\n")
}

func (d IODevice) readStringArray(file string) []string {
	return strings.Split(d.readString(file), " ")
}

func (d IODevice) readStringSelector(file string) string {
	return d.readString(file)
}
