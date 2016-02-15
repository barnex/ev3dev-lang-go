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
	"time"
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
	// 360
}

func ExampleMotor_DriverName() {
	medium, _ := OpenMotor("outA")
	fmt.Println(medium.DriverName())
	large, _ := OpenMotor("outB")
	fmt.Println(large.DriverName())

	//Output:
	// lego-ev3-m-motor
	// lego-ev3-l-motor
}

func ExampleMotor_DutyCycle() {
	m, _ := OpenMotor("")
	fmt.Println(m.DutyCycle())

	//Output:
	// 0
}

func ExampleMotor_DutyCycleSP() {
	m, _ := OpenMotor("outA")
	m.SetDutyCycleSP(100)
	fmt.Println(m.DutyCycleSP())

	//Output:
	// 100
}

func ExampleMotor_SetDutyCycleSP() {
	m, _ := OpenMotor("outA")
	m.SetDutyCycleSP(100)
	fmt.Println(m.DutyCycleSP())

	//Output:
	// 100
}

func ExampleMotor_EncoderPolarity() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.EncoderPolarity())

	//Output:
	// normal
}

func ExampleMotor_SetEncoderPolarity() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.EncoderPolarity())
	m.SetEncoderPolarity("inversed")
	fmt.Println(m.EncoderPolarity())

	//Output:
	// normal
	// inversed
}

func ExampleMotor_Polarity() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.Polarity())

	//Output:
	// normal
}

func ExampleMotor_SetPolarity() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.Polarity())
	m.SetPolarity("inversed")
	fmt.Println(m.Polarity())

	//Output:
	// normal
	// inversed
}

func ExampleMotor_Position() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.Position())

	//Output:
	// 0
}

func ExampleMotor_SetPosition() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.Position())
	m.SetPosition(200)
	m.SetSpeedSP(10)
	m.SetCommand("run-to-abs-pos")
	fmt.Println(m.Position())

	//Output:
	// 0
	// 200
}

func ExampleMotor_PositionD() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.PositionD())

	//Output:
	// 0
}
func ExampleMotor_PositionI() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.PositionI())

	//Output:
	// 0
}
func ExampleMotor_PositionP() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.PositionP())

	//Output:
	// 0

}
func ExampleMotor_PositionSP() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.PositionSP())

	//Output:
	// 0
}
func ExampleMotor_RampDownSP() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.RampDownSP())

	//Output:
	// 0

}
func ExampleMotor_RampUpSP() {
	m, _ := OpenMotor("outA")
	fmt.Println(m.RampUpSP())

	//Output:
	// 0

}

func ExampleMotor_SetCommand() {
	m, _ := OpenMotor("outA")

	m.SetDutyCycleSP(50)
	fmt.Println("start")
	m.SetCommand("run-forever")
	time.Sleep(1 * time.Second)
	m.SetCommand("stop")
	fmt.Println("stop")

	//Output:
	// start
	// stop
}

func ExampleMotor_SetRampDownSP() {}

func ExampleMotor_SetRampUpSP() {}

func ExampleMotor_SetSpeedRegulationD() {}

func ExampleMotor_SetSpeedRegulationEnabled() {}

func ExampleMotor_SetSpeedRegulationI() {}

func ExampleMotor_SetSpeedRegulationP() {}

func ExampleMotor_SetSpeedSP() {}

func ExampleMotor_StopCommand() {
	m, _ := OpenMotor("outB")
	fmt.Println(m.StopCommand())

	//Output:
	// coast
}

func ExampleMotor_SetStopCommand() {
	m, _ := OpenMotor("outB")
	fmt.Println(m.StopCommand())
	m.SetStopCommand("brake")
	fmt.Println(m.StopCommand())
	m.SetStopCommand("hold")
	fmt.Println(m.StopCommand())

	//Output:
	// coast
	// brake
	// hold
}

func ExampleMotor_SetTimeSP() {}

func ExampleMotor_Speed() {}

func ExampleMotor_SpeedRegulationD() {}

func ExampleMotor_SpeedRegulationEnabled() {}

func ExampleMotor_SpeedRegulationI() {}

func ExampleMotor_SpeedRegulationP() {}

func ExampleMotor_SpeedSP() {}

func ExampleMotor_State() {}

func ExampleMotor_StopCommands() {
	m, _ := OpenMotor("outB")
	fmt.Println(m.StopCommands)

	//Output:
	// [coast brake hold]
}

func ExampleMotor_TimeSP() {}
