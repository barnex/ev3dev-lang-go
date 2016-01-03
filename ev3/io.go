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

func(d IODevice)write(file string, x interface{}){
	f, err := os.OpenFile(path.Join(d.path, file), os.O_WRONLY, 0666)
	if err != nil{panic (err)}
	defer f.Close()
	fmt.Println(f, x)
}

func(d IODevice)read(file string) []byte{
	f, err := os.OpenFile(path.Join(d.path, file), os.O_RDONLY, 0666)
	if err != nil{panic (err)}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil{panic(err)}
	return bytes
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
	return string(d.read(file))
}

func (d IODevice) readStringArray(file string) []string {
	return strings.Split(d.readString(file), " ")
}

func (d IODevice) readStringSelector(file string) string {
	return d.readString(file)
}
