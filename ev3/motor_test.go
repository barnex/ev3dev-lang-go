package ev3

/*
Hardware setup:

	outA: medium motor
	outB: large motor
	outC: large motor
	outD: -

	in1: -
	in2: touch sensor
	in3: color sensor
	in4: IR sensor

*/

import (
	"fmt"
	"log"
)

func init() {
	//SysFS = "testdata/"
}

func ExampleOpenMotor() {
	// Opens the first motor found
	motor1, err := OpenMotor("")
	fmt.Println(motor1, err)

	motorA, err := OpenMotor("outA")
	fmt.Println(motorA, err)

	// Not connected: returns error
	motorD, err := OpenMotor("outD")
	fmt.Println(motorD, err)

	//Output:
	// /sys/class/tacho-motor/motor0 <nil>
	// /sys/class/tacho-motor/motor0 <nil>
	// <nil> open tacho-motor:outD: port not found in [motor0 motor1 motor2]
}

func ExampleMotor_Address() {
	motor1, err := OpenMotor("")
	fmt.Println(motor1.Address(), err)

	motorC, err := OpenMotor("outC")
	fmt.Println(motorC.Address(), err)

	//Output:
	// outA <nil>
	// outC <nil>
}

func ExampleMotor_Commands() {
	m, err := OpenMotor("outA")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.Commands())

	//Output:
	// [run-forever run-to-abs-pos run-to-rel-pos run-timed run-direct stop reset]
}

func ExampleMotor_CountPerRot() {

	m, err := OpenMotor("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(m.CountPerRot())

	//Output:
	// 0
}

func ExampleMotor_DriverName()                {}
func ExampleMotor_DutyCycle()                 {}
func ExampleMotor_DutyCycleSP()               {}
func ExampleMotor_EncoderPolarity()           {}
func ExampleMotor_Polarity()                  {}
func ExampleMotor_Position()                  {}
func ExampleMotor_PositionD()                 {}
func ExampleMotor_PositionI()                 {}
func ExampleMotor_PositionP()                 {}
func ExampleMotor_PositionSP()                {}
func ExampleMotor_RampDownSP()                {}
func ExampleMotor_RampUpSP()                  {}
func ExampleMotor_SetCommand()                {}
func ExampleMotor_SetDutyCycleSP()            {}
func ExampleMotor_SetEncoderPolarity()        {}
func ExampleMotor_SetPolarity()               {}
func ExampleMotor_SetPosition()               {}
func ExampleMotor_SetPositionD()              {}
func ExampleMotor_SetPositionI()              {}
func ExampleMotor_SetPositionP()              {}
func ExampleMotor_SetPositionSP()             {}
func ExampleMotor_SetRampDownSP()             {}
func ExampleMotor_SetRampUpSP()               {}
func ExampleMotor_SetSpeedRegulationD()       {}
func ExampleMotor_SetSpeedRegulationEnabled() {}
func ExampleMotor_SetSpeedRegulationI()       {}
func ExampleMotor_SetSpeedRegulationP()       {}
func ExampleMotor_SetSpeedSP()                {}
func ExampleMotor_SetStopCommand()            {}
func ExampleMotor_SetTimeSP()                 {}
func ExampleMotor_Speed()                     {}
func ExampleMotor_SpeedRegulationD()          {}
func ExampleMotor_SpeedRegulationEnabled()    {}
func ExampleMotor_SpeedRegulationI()          {}
func ExampleMotor_SpeedRegulationP()          {}
func ExampleMotor_SpeedSP()                   {}
func ExampleMotor_State()                     {}
func ExampleMotor_StopCommand()               {}
func ExampleMotor_StopCommands()              {}
func ExampleMotor_TimeSP()                    {}
