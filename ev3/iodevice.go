package ev3

type IODevice struct {
	port string
}

func OpenIODevice(sysClass, devNameConvention, port string) (IODevice, error) {
	panic("not implemented")
}

type Selector string
