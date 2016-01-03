package ev3

var SysFS = "/sys/class/"

type IODevice struct {
	port string
}

func OpenIODevice(sysClass, devNameConvention, port string) (IODevice, error) {
	panic("not implemented")
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
