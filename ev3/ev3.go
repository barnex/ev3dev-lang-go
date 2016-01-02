// UNDER CONSTRUCTION, NOTHING TO SEE HERE
//
// Package ev3 provides a Go interface for LEGO Mindstorms EV3 running ev3dev
// (http://www.ev3dev.org).
// This file is auto-generated for the spec defined at http://github.com/ev3dev/ev3dev-lang, version 0.9.3-pre, revision 2.
// Supported kernel: v3.16.7-ckt16-7-ev3dev-ev3
package ev3

// Provides a generic button reading mechanism that can be adapted
// to platform specific implementations. Each platform's specific
// button capabilites are enumerated in the 'platforms' section
// of this specification
type Button struct {
	IODevice
}

func OpenButton(port string) (*Button, error) {

	io, err := OpenIODevice("<no value>", "<no value>", port)
	if err != nil {
		return nil, err
	}
	return &Button{IODevice: io}, nil

}

// The DC motor class provides a uniform interface for using regular DC motors
// with no fancy controls or feedback. This includes LEGO MINDSTORMS RCX motors
// and LEGO Power Functions motors.
// http://www.ev3dev.org/docs/drivers/dc-motor-class/
type DcMotor struct {
	IODevice
}

func OpenDcMotor(port string) (*DcMotor, error) {

	io, err := OpenIODevice("dc-motor", "motor{0}", port)
	if err != nil {
		return nil, err
	}
	return &DcMotor{IODevice: io}, nil

}

// Sets the command for the motor. Possible values are `run-forever`, `run-timed` and
// `stop`. Not all commands may be supported, so be sure to check the contents
// of the `commands` attribute.
func (d *DcMotor) SetCommand(x string) {
	d.writeString("command", x)
}

// Returns a list of commands supported by the motor
// controller.
func (d *DcMotor) Commands() []string {
	return d.readStringArray("commands")
}

// Returns the name of the motor driver that loaded this device. See the list
// of [supported devices] for a list of drivers.
func (d *DcMotor) DriverName() string {
	return d.readString("driver_name")
}

// Shows the current duty cycle of the PWM signal sent to the motor. Values
// are -100 to 100 (-100% to 100%).
func (d *DcMotor) DutyCycle() int {
	return d.readInt("duty_cycle")
}

// Writing sets the duty cycle setpoint of the PWM signal sent to the motor.
// Valid values are -100 to 100 (-100% to 100%). Reading returns the current
// setpoint.
func (d *DcMotor) SetDutyCycleSp(x int) {
	d.writeInt("duty_cycle_sp", x)
}

// Writing sets the duty cycle setpoint of the PWM signal sent to the motor.
// Valid values are -100 to 100 (-100% to 100%). Reading returns the current
// setpoint.
func (d *DcMotor) DutyCycleSp() int {
	return d.readInt("duty_cycle_sp")
}

// Sets the polarity of the motor. Valid values are `normal` and `inversed`.
func (d *DcMotor) SetPolarity(x string) {
	d.writeString("polarity", x)
}

// Sets the polarity of the motor. Valid values are `normal` and `inversed`.
func (d *DcMotor) Polarity() string {
	return d.readString("polarity")
}

// Returns the name of the port that this motor is connected to.
func (d *DcMotor) Address() string {
	return d.readString("address")
}

// Sets the time in milliseconds that it take the motor to ramp down from 100%
// to 0%. Valid values are 0 to 10000 (10 seconds). Default is 0.
func (d *DcMotor) SetRampDownSp(x int) {
	d.writeInt("ramp_down_sp", x)
}

// Sets the time in milliseconds that it take the motor to ramp down from 100%
// to 0%. Valid values are 0 to 10000 (10 seconds). Default is 0.
func (d *DcMotor) RampDownSp() int {
	return d.readInt("ramp_down_sp")
}

// Sets the time in milliseconds that it take the motor to up ramp from 0% to
// 100%. Valid values are 0 to 10000 (10 seconds). Default is 0.
func (d *DcMotor) SetRampUpSp(x int) {
	d.writeInt("ramp_up_sp", x)
}

// Sets the time in milliseconds that it take the motor to up ramp from 0% to
// 100%. Valid values are 0 to 10000 (10 seconds). Default is 0.
func (d *DcMotor) RampUpSp() int {
	return d.readInt("ramp_up_sp")
}

// Gets a list of flags indicating the motor status. Possible
// flags are `running` and `ramping`. `running` indicates that the motor is
// powered. `ramping` indicates that the motor has not yet reached the
// `duty_cycle_sp`.
func (d *DcMotor) State() []string {
	return d.readStringArray("state")
}

// Sets the stop command that will be used when the motor stops. Read
// `stop_commands` to get the list of valid values.
func (d *DcMotor) SetStopCommand(x string) {
	d.writeString("stop_command", x)
}

// Gets a list of stop commands. Valid values are `coast`
// and `brake`.
func (d *DcMotor) StopCommands() []string {
	return d.readStringArray("stop_commands")
}

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
func (d *DcMotor) SetTimeSp(x int) {
	d.writeInt("time_sp", x)
}

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
func (d *DcMotor) TimeSp() int {
	return d.readInt("time_sp")
}

// A generic interface to control I2C-type EV3 sensors.
type I2cSensor struct {
	Sensor
}

func OpenI2cSensor(port string) (*I2cSensor, error) {

	super, err := OpenSensor(port)
	if err != nil {
		return nil, err
	}
	return &I2cSensor{*super}, nil

}

// Returns the firmware version of the sensor if available. Currently only
// I2C/NXT sensors support this.
func (i *I2cSensor) FwVersion() string {
	return i.readString("fw_version")
}

// Returns the polling period of the sensor in milliseconds. Writing sets the
// polling period. Setting to 0 disables polling. Minimum value is hard
// coded as 50 msec. Returns -EOPNOTSUPP if changing polling is not supported.
// Currently only I2C/NXT sensors support changing the polling period.
func (i *I2cSensor) SetPollMs(x int) {
	i.writeInt("poll_ms", x)
}

// Returns the polling period of the sensor in milliseconds. Writing sets the
// polling period. Setting to 0 disables polling. Minimum value is hard
// coded as 50 msec. Returns -EOPNOTSUPP if changing polling is not supported.
// Currently only I2C/NXT sensors support changing the polling period.
func (i *I2cSensor) PollMs() int {
	return i.readInt("poll_ms")
}

// EV3 large servo motor
type LargeMotor struct {
	Motor
}

func OpenLargeMotor(port string) (*LargeMotor, error) {

	super, err := OpenMotor(port)
	if err != nil {
		return nil, err
	}
	return &LargeMotor{*super}, nil

}

// Any device controlled by the generic LED driver.
// See https://www.kernel.org/doc/Documentation/leds/leds-class.txt
// for more details.
type Led struct {
	IODevice
}

func OpenLed(port string) (*Led, error) {

	io, err := OpenIODevice("leds", "<no value>", port)
	if err != nil {
		return nil, err
	}
	return &Led{IODevice: io}, nil

}

// Returns the maximum allowable brightness value.
func (l *Led) MaxBrightness() int {
	return l.readInt("max_brightness")
}

// Sets the brightness level. Possible values are from 0 to `max_brightness`.
func (l *Led) SetBrightness(x int) {
	l.writeInt("brightness", x)
}

// Sets the brightness level. Possible values are from 0 to `max_brightness`.
func (l *Led) Brightness() int {
	return l.readInt("brightness")
}

// Returns a list of available triggers.
func (l *Led) Triggers() []string {
	return l.readStringArray("triggers")
}

// Sets the led trigger. A trigger
// is a kernel based source of led events. Triggers can either be simple or
// complex. A simple trigger isn't configurable and is designed to slot into
// existing subsystems with minimal additional code. Examples are the `ide-disk` and
// `nand-disk` triggers.
//
// Complex triggers whilst available to all LEDs have LED specific
// parameters and work on a per LED basis. The `timer` trigger is an example.
// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` and `off` time can
// be specified via `delay_{on,off}` attributes in milliseconds.
// You can change the brightness value of a LED independently of the timer
// trigger. However, if you set the brightness value to 0 it will
// also disable the `timer` trigger.
func (l *Led) SetTrigger(x string) {
	l.writeStringSelector("trigger", x)
}

// Sets the led trigger. A trigger
// is a kernel based source of led events. Triggers can either be simple or
// complex. A simple trigger isn't configurable and is designed to slot into
// existing subsystems with minimal additional code. Examples are the `ide-disk` and
// `nand-disk` triggers.
//
// Complex triggers whilst available to all LEDs have LED specific
// parameters and work on a per LED basis. The `timer` trigger is an example.
// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` and `off` time can
// be specified via `delay_{on,off}` attributes in milliseconds.
// You can change the brightness value of a LED independently of the timer
// trigger. However, if you set the brightness value to 0 it will
// also disable the `timer` trigger.
func (l *Led) Trigger() string {
	return l.readStringSelector("trigger")
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` time can
// be specified via `delay_on` attribute in milliseconds.
func (l *Led) SetDelayOn(x int) {
	l.writeInt("delay_on", x)
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `on` time can
// be specified via `delay_on` attribute in milliseconds.
func (l *Led) DelayOn() int {
	return l.readInt("delay_on")
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `off` time can
// be specified via `delay_off` attribute in milliseconds.
func (l *Led) SetDelayOff(x int) {
	l.writeInt("delay_off", x)
}

// The `timer` trigger will periodically change the LED brightness between
// 0 and the current brightness setting. The `off` time can
// be specified via `delay_off` attribute in milliseconds.
func (l *Led) DelayOff() int {
	return l.readInt("delay_off")
}

// The `lego-port` class provides an interface for working with input and
// output ports that are compatible with LEGO MINDSTORMS RCX/NXT/EV3, LEGO
// WeDo and LEGO Power Functions sensors and motors. Supported devices include
// the LEGO MINDSTORMS EV3 Intelligent Brick, the LEGO WeDo USB hub and
// various sensor multiplexers from 3rd party manufacturers.
//
// Some types of ports may have multiple modes of operation. For example, the
// input ports on the EV3 brick can communicate with sensors using UART, I2C
// or analog validate signals - but not all at the same time. Therefore there
// are multiple modes available to connect to the different types of sensors.
//
// In most cases, ports are able to automatically detect what type of sensor
// or motor is connected. In some cases though, this must be manually specified
// using the `mode` and `set_device` attributes. The `mode` attribute affects
// how the port communicates with the connected device. For example the input
// ports on the EV3 brick can communicate using UART, I2C or analog voltages,
// but not all at the same time, so the mode must be set to the one that is
// appropriate for the connected sensor. The `set_device` attribute is used to
// specify the exact type of sensor that is connected. Note: the mode must be
// correctly set before setting the sensor type.
//
// Ports can be found at `/sys/class/lego-port/port<N>` where `<N>` is
// incremented each time a new port is registered. Note: The number is not
// related to the actual port at all - use the `address` attribute to find
// a specific port.
type LegoPort struct {
	IODevice
}

func OpenLegoPort(port string) (*LegoPort, error) {

	io, err := OpenIODevice("lego_port", "<no value>", port)
	if err != nil {
		return nil, err
	}
	return &LegoPort{IODevice: io}, nil

}

// Returns the name of the driver that loaded this device. You can find the
// complete list of drivers in the [list of port drivers].
func (l *LegoPort) DriverName() string {
	return l.readString("driver_name")
}

// Returns a list of the available modes of the port.
func (l *LegoPort) Modes() []string {
	return l.readStringArray("modes")
}

// Reading returns the currently selected mode. Writing sets the mode.
// Generally speaking when the mode changes any sensor or motor devices
// associated with the port will be removed new ones loaded, however this
// this will depend on the individual driver implementing this class.
func (l *LegoPort) SetMode(x string) {
	l.writeString("mode", x)
}

// Reading returns the currently selected mode. Writing sets the mode.
// Generally speaking when the mode changes any sensor or motor devices
// associated with the port will be removed new ones loaded, however this
// this will depend on the individual driver implementing this class.
func (l *LegoPort) Mode() string {
	return l.readString("mode")
}

// Returns the name of the port. See individual driver documentation for
// the name that will be returned.
func (l *LegoPort) Address() string {
	return l.readString("address")
}

// For modes that support it, writing the name of a driver will cause a new
// device to be registered for that driver and attached to this port. For
// example, since NXT/Analog sensors cannot be auto-detected, you must use
// this attribute to load the correct driver. Returns -EOPNOTSUPP if setting a
// device is not supported.
func (l *LegoPort) SetSetDevice(x string) {
	l.writeString("set_device", x)
}

// In most cases, reading status will return the same value as `mode`. In
// cases where there is an `auto` mode additional values may be returned,
// such as `no-device` or `error`. See individual port driver documentation
// for the full list of possible values.
func (l *LegoPort) Status() string {
	return l.readString("status")
}

// EV3 medium servo motor
type MediumMotor struct {
	Motor
}

func OpenMediumMotor(port string) (*MediumMotor, error) {

	super, err := OpenMotor(port)
	if err != nil {
		return nil, err
	}
	return &MediumMotor{*super}, nil

}

// The motor class provides a uniform interface for using motors with
// positional and directional feedback such as the EV3 and NXT motors.
// This feedback allows for precise control of the motors. This is the
// most common type of motor, so we just call it `motor`.
// http://www.ev3dev.org/docs/drivers/tacho-motor-class/
type Motor struct {
	IODevice
}

func OpenMotor(port string) (*Motor, error) {

	io, err := OpenIODevice("tacho-motor", "motor{0}", port)
	if err != nil {
		return nil, err
	}
	return &Motor{IODevice: io}, nil

}

// Sends a command to the motor controller. See `commands` for a list of
// possible values.
func (m *Motor) SetCommand(x string) {
	m.writeString("command", x)
}

// Returns a list of commands that are supported by the motor
// controller. Possible values are `run-forever`, `run-to-abs-pos`, `run-to-rel-pos`,
// `run-timed`, `run-direct`, `stop` and `reset`. Not all commands may be supported.
//
//  - `run-forever` will cause the motor to run until another command is sent.
//  - `run-to-abs-pos` will run to an absolute position specified by `position_sp`
//   and then stop using the command specified in `stop_command`.
//  - `run-to-rel-pos` will run to a position relative to the current `position` value.
//   The new position will be current `position` + `position_sp`. When the new
//   position is reached, the motor will stop using the command specified by `stop_command`.
//  - `run-timed` will run the motor for the amount of time specified in `time_sp`
//   and then stop the motor using the command specified by `stop_command`.
//  - `run-direct` will run the motor at the duty cycle specified by `duty_cycle_sp`.
//   Unlike other run commands, changing `duty_cycle_sp` while running *will*
//   take effect immediately.
//  - `stop` will stop any of the run commands before they are complete using the
//   command specified by `stop_command`.
//  - `reset` will reset all of the motor parameter attributes to their default value.
//   This will also have the effect of stopping the motor.
func (m *Motor) Commands() []string {
	return m.readStringArray("commands")
}

// Returns the number of tacho counts in one rotation of the motor. Tacho counts
// are used by the position and speed attributes, so you can use this value
// to convert rotations or degrees to tacho counts. In the case of linear
// actuators, the units here will be counts per centimeter.
func (m *Motor) CountPerRot() int {
	return m.readInt("count_per_rot")
}

// Returns the name of the driver that provides this tacho motor device.
func (m *Motor) DriverName() string {
	return m.readString("driver_name")
}

// Returns the current duty cycle of the motor. Units are percent. Values
// are -100 to 100.
func (m *Motor) DutyCycle() int {
	return m.readInt("duty_cycle")
}

// Writing sets the duty cycle setpoint. Reading returns the current value.
// Units are in percent. Valid values are -100 to 100. A negative value causes
// the motor to rotate in reverse. This value is only used when `speed_regulation`
// is off.
func (m *Motor) SetDutyCycleSp(x int) {
	m.writeInt("duty_cycle_sp", x)
}

// Writing sets the duty cycle setpoint. Reading returns the current value.
// Units are in percent. Valid values are -100 to 100. A negative value causes
// the motor to rotate in reverse. This value is only used when `speed_regulation`
// is off.
func (m *Motor) DutyCycleSp() int {
	return m.readInt("duty_cycle_sp")
}

// Sets the polarity of the rotary encoder. This is an advanced feature to all
// use of motors that send inversed encoder signals to the EV3. This should
// be set correctly by the driver of a device. It You only need to change this
// value if you are using a unsupported device. Valid values are `normal` and
// `inversed`.
func (m *Motor) SetEncoderPolarity(x string) {
	m.writeString("encoder_polarity", x)
}

// Sets the polarity of the rotary encoder. This is an advanced feature to all
// use of motors that send inversed encoder signals to the EV3. This should
// be set correctly by the driver of a device. It You only need to change this
// value if you are using a unsupported device. Valid values are `normal` and
// `inversed`.
func (m *Motor) EncoderPolarity() string {
	return m.readString("encoder_polarity")
}

// Sets the polarity of the motor. With `normal` polarity, a positive duty
// cycle will cause the motor to rotate clockwise. With `inversed` polarity,
// a positive duty cycle will cause the motor to rotate counter-clockwise.
// Valid values are `normal` and `inversed`.
func (m *Motor) SetPolarity(x string) {
	m.writeString("polarity", x)
}

// Sets the polarity of the motor. With `normal` polarity, a positive duty
// cycle will cause the motor to rotate clockwise. With `inversed` polarity,
// a positive duty cycle will cause the motor to rotate counter-clockwise.
// Valid values are `normal` and `inversed`.
func (m *Motor) Polarity() string {
	return m.readString("polarity")
}

// Returns the name of the port that this motor is connected to.
func (m *Motor) Address() string {
	return m.readString("address")
}

// Returns the current position of the motor in pulses of the rotary
// encoder. When the motor rotates clockwise, the position will increase.
// Likewise, rotating counter-clockwise causes the position to decrease.
// Writing will set the position to that value.
func (m *Motor) SetPosition(x int) {
	m.writeInt("position", x)
}

// Returns the current position of the motor in pulses of the rotary
// encoder. When the motor rotates clockwise, the position will increase.
// Likewise, rotating counter-clockwise causes the position to decrease.
// Writing will set the position to that value.
func (m *Motor) Position() int {
	return m.readInt("position")
}

// The proportional constant for the position PID.
func (m *Motor) SetHoldPidKp(x int) {
	m.writeInt("hold_pid/Kp", x)
}

// The proportional constant for the position PID.
func (m *Motor) HoldPidKp() int {
	return m.readInt("hold_pid/Kp")
}

// The integral constant for the position PID.
func (m *Motor) SetHoldPidKi(x int) {
	m.writeInt("hold_pid/Ki", x)
}

// The integral constant for the position PID.
func (m *Motor) HoldPidKi() int {
	return m.readInt("hold_pid/Ki")
}

// The derivative constant for the position PID.
func (m *Motor) SetHoldPidKd(x int) {
	m.writeInt("hold_pid/Kd", x)
}

// The derivative constant for the position PID.
func (m *Motor) HoldPidKd() int {
	return m.readInt("hold_pid/Kd")
}

// Writing specifies the target position for the `run-to-abs-pos` and `run-to-rel-pos`
// commands. Reading returns the current value. Units are in tacho counts. You
// can use the value returned by `counts_per_rot` to convert tacho counts to/from
// rotations or degrees.
func (m *Motor) SetPositionSp(x int) {
	m.writeInt("position_sp", x)
}

// Writing specifies the target position for the `run-to-abs-pos` and `run-to-rel-pos`
// commands. Reading returns the current value. Units are in tacho counts. You
// can use the value returned by `counts_per_rot` to convert tacho counts to/from
// rotations or degrees.
func (m *Motor) PositionSp() int {
	return m.readInt("position_sp")
}

// Returns the current motor speed in tacho counts per second. Not, this is
// not necessarily degrees (although it is for LEGO motors). Use the `count_per_rot`
// attribute to convert this value to RPM or deg/sec.
func (m *Motor) Speed() int {
	return m.readInt("speed")
}

// Writing sets the target speed in tacho counts per second used when `speed_regulation`
// is on. Reading returns the current value.  Use the `count_per_rot` attribute
// to convert RPM or deg/sec to tacho counts per second.
func (m *Motor) SetSpeedSp(x int) {
	m.writeInt("speed_sp", x)
}

// Writing sets the target speed in tacho counts per second used when `speed_regulation`
// is on. Reading returns the current value.  Use the `count_per_rot` attribute
// to convert RPM or deg/sec to tacho counts per second.
func (m *Motor) SpeedSp() int {
	return m.readInt("speed_sp")
}

// Writing sets the ramp up setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 0 to 100% duty cycle over the span of this setpoint
// when starting the motor. If the maximum duty cycle is limited by `duty_cycle_sp`
// or speed regulation, the actual ramp time duration will be less than the setpoint.
func (m *Motor) SetRampUpSp(x int) {
	m.writeInt("ramp_up_sp", x)
}

// Writing sets the ramp up setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 0 to 100% duty cycle over the span of this setpoint
// when starting the motor. If the maximum duty cycle is limited by `duty_cycle_sp`
// or speed regulation, the actual ramp time duration will be less than the setpoint.
func (m *Motor) RampUpSp() int {
	return m.readInt("ramp_up_sp")
}

// Writing sets the ramp down setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 100% duty cycle down to 0 over the span of this setpoint
// when stopping the motor. If the starting duty cycle is less than 100%, the
// ramp time duration will be less than the full span of the setpoint.
func (m *Motor) SetRampDownSp(x int) {
	m.writeInt("ramp_down_sp", x)
}

// Writing sets the ramp down setpoint. Reading returns the current value. Units
// are in milliseconds. When set to a value > 0, the motor will ramp the power
// sent to the motor from 100% duty cycle down to 0 over the span of this setpoint
// when stopping the motor. If the starting duty cycle is less than 100%, the
// ramp time duration will be less than the full span of the setpoint.
func (m *Motor) RampDownSp() int {
	return m.readInt("ramp_down_sp")
}

// Turns speed regulation on or off. If speed regulation is on, the motor
// controller will vary the power supplied to the motor to try to maintain the
// speed specified in `speed_sp`. If speed regulation is off, the controller
// will use the power specified in `duty_cycle_sp`. Valid values are `on` and
// `off`.
func (m *Motor) SetSpeedRegulation(x string) {
	m.writeString("speed_regulation", x)
}

// Turns speed regulation on or off. If speed regulation is on, the motor
// controller will vary the power supplied to the motor to try to maintain the
// speed specified in `speed_sp`. If speed regulation is off, the controller
// will use the power specified in `duty_cycle_sp`. Valid values are `on` and
// `off`.
func (m *Motor) SpeedRegulation() string {
	return m.readString("speed_regulation")
}

// The proportional constant for the speed regulation PID.
func (m *Motor) SetSpeedPidKp(x int) {
	m.writeInt("speed_pid/Kp", x)
}

// The proportional constant for the speed regulation PID.
func (m *Motor) SpeedPidKp() int {
	return m.readInt("speed_pid/Kp")
}

// The integral constant for the speed regulation PID.
func (m *Motor) SetSpeedPidKi(x int) {
	m.writeInt("speed_pid/Ki", x)
}

// The integral constant for the speed regulation PID.
func (m *Motor) SpeedPidKi() int {
	return m.readInt("speed_pid/Ki")
}

// The derivative constant for the speed regulation PID.
func (m *Motor) SetSpeedPidKd(x int) {
	m.writeInt("speed_pid/Kd", x)
}

// The derivative constant for the speed regulation PID.
func (m *Motor) SpeedPidKd() int {
	return m.readInt("speed_pid/Kd")
}

// Reading returns a list of state flags. Possible flags are
// `running`, `ramping` `holding` and `stalled`.
func (m *Motor) State() []string {
	return m.readStringArray("state")
}

// Reading returns the current stop command. Writing sets the stop command.
// The value determines the motors behavior when `command` is set to `stop`.
// Also, it determines the motors behavior when a run command completes. See
// `stop_commands` for a list of possible values.
func (m *Motor) SetStopCommand(x string) {
	m.writeString("stop_command", x)
}

// Reading returns the current stop command. Writing sets the stop command.
// The value determines the motors behavior when `command` is set to `stop`.
// Also, it determines the motors behavior when a run command completes. See
// `stop_commands` for a list of possible values.
func (m *Motor) StopCommand() string {
	return m.readString("stop_command")
}

// Returns a list of stop modes supported by the motor controller.
// Possible values are `coast`, `brake` and `hold`. `coast` means that power will
// be removed from the motor and it will freely coast to a stop. `brake` means
// that power will be removed from the motor and a passive electrical load will
// be placed on the motor. This is usually done by shorting the motor terminals
// together. This load will absorb the energy from the rotation of the motors and
// cause the motor to stop more quickly than coasting. `hold` does not remove
// power from the motor. Instead it actively try to hold the motor at the current
// position. If an external force tries to turn the motor, the motor will 'push
// back' to maintain its position.
func (m *Motor) StopCommands() []string {
	return m.readStringArray("stop_commands")
}

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
func (m *Motor) SetTimeSp(x int) {
	m.writeInt("time_sp", x)
}

// Writing specifies the amount of time the motor will run when using the
// `run-timed` command. Reading returns the current value. Units are in
// milliseconds.
func (m *Motor) TimeSp() int {
	return m.readInt("time_sp")
}

// A generic interface to read data from the system's power_supply class.
// Uses the built-in legoev3-battery if none is specified.
type PowerSupply struct {
	IODevice
}

func OpenPowerSupply(port string) (*PowerSupply, error) {

	io, err := OpenIODevice("power_supply", "<no value>", port)
	if err != nil {
		return nil, err
	}
	return &PowerSupply{IODevice: io}, nil

}

// The measured current that the battery is supplying (in microamps)
func (p *PowerSupply) CurrentNow() int {
	return p.readInt("current_now")
}

// The measured voltage that the battery is supplying (in microvolts)
func (p *PowerSupply) VoltageNow() int {
	return p.readInt("voltage_now")
}

func (p *PowerSupply) VoltageMaxDesign() int {
	return p.readInt("voltage_max_design")
}

func (p *PowerSupply) VoltageMinDesign() int {
	return p.readInt("voltage_min_design")
}

func (p *PowerSupply) Technology() string {
	return p.readString("technology")
}

func (p *PowerSupply) Type() string {
	return p.readString("type")
}

// The sensor class provides a uniform interface for using most of the
// sensors available for the EV3. The various underlying device drivers will
// create a `lego-sensor` device for interacting with the sensors.
//
// Sensors are primarily controlled by setting the `mode` and monitored by
// reading the `value<N>` attributes. Values can be converted to floating point
// if needed by `value<N>` / 10.0 ^ `decimals`.
//
// Since the name of the `sensor<N>` device node does not correspond to the port
// that a sensor is plugged in to, you must look at the `address` attribute if
// you need to know which port a sensor is plugged in to. However, if you don't
// have more than one sensor of each type, you can just look for a matching
// `driver_name`. Then it will not matter which port a sensor is plugged in to - your
// program will still work.
// http://www.ev3dev.org/docs/drivers/lego-sensor-class/
type Sensor struct {
	IODevice
}

func OpenSensor(port string) (*Sensor, error) {

	io, err := OpenIODevice("lego-sensor", "sensor{0}", port)
	if err != nil {
		return nil, err
	}
	return &Sensor{IODevice: io}, nil

}

// Sends a command to the sensor.
func (s *Sensor) SetCommand(x string) {
	s.writeString("command", x)
}

// Returns a list of the valid commands for the sensor.
// Returns -EOPNOTSUPP if no commands are supported.
func (s *Sensor) Commands() []string {
	return s.readStringArray("commands")
}

// Returns the number of decimal places for the values in the `value<N>`
// attributes of the current mode.
func (s *Sensor) Decimals() int {
	return s.readInt("decimals")
}

// Returns the name of the sensor device/driver. See the list of [supported
// sensors] for a complete list of drivers.
func (s *Sensor) DriverName() string {
	return s.readString("driver_name")
}

// Returns the current mode. Writing one of the values returned by `modes`
// sets the sensor to that mode.
func (s *Sensor) SetMode(x string) {
	s.writeString("mode", x)
}

// Returns the current mode. Writing one of the values returned by `modes`
// sets the sensor to that mode.
func (s *Sensor) Mode() string {
	return s.readString("mode")
}

// Returns a list of the valid modes for the sensor.
func (s *Sensor) Modes() []string {
	return s.readStringArray("modes")
}

// Returns the number of `value<N>` attributes that will return a valid value
// for the current mode.
func (s *Sensor) NumValues() int {
	return s.readInt("num_values")
}

// Returns the name of the port that the sensor is connected to, e.g. `ev3:in1`.
// I2C sensors also include the I2C address (decimal), e.g. `ev3:in1:i2c8`.
func (s *Sensor) Address() string {
	return s.readString("address")
}

// Returns the units of the measured value for the current mode. May return
// empty string
func (s *Sensor) Units() string {
	return s.readString("units")
}

// The servo motor class provides a uniform interface for using hobby type
// servo motors.
// http://www.ev3dev.org/docs/drivers/servo-motor-class/
type ServoMotor struct {
	IODevice
}

func OpenServoMotor(port string) (*ServoMotor, error) {

	io, err := OpenIODevice("servo-motor", "motor{0}", port)
	if err != nil {
		return nil, err
	}
	return &ServoMotor{IODevice: io}, nil

}

// Sets the command for the servo. Valid values are `run` and `float`. Setting
// to `run` will cause the servo to be driven to the position_sp set in the
// `position_sp` attribute. Setting to `float` will remove power from the motor.
func (s *ServoMotor) SetCommand(x string) {
	s.writeString("command", x)
}

// Returns the name of the motor driver that loaded this device. See the list
// of [supported devices] for a list of drivers.
func (s *ServoMotor) DriverName() string {
	return s.readString("driver_name")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the maximum (clockwise) position_sp. Default value is 2400.
// Valid values are 2300 to 2700. You must write to the position_sp attribute for
// changes to this attribute to take effect.
func (s *ServoMotor) SetMaxPulseSp(x int) {
	s.writeInt("max_pulse_sp", x)
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the maximum (clockwise) position_sp. Default value is 2400.
// Valid values are 2300 to 2700. You must write to the position_sp attribute for
// changes to this attribute to take effect.
func (s *ServoMotor) MaxPulseSp() int {
	return s.readInt("max_pulse_sp")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the mid position_sp. Default value is 1500. Valid
// values are 1300 to 1700. For example, on a 180 degree servo, this would be
// 90 degrees. On continuous rotation servo, this is the 'neutral' position_sp
// where the motor does not turn. You must write to the position_sp attribute for
// changes to this attribute to take effect.
func (s *ServoMotor) SetMidPulseSp(x int) {
	s.writeInt("mid_pulse_sp", x)
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the mid position_sp. Default value is 1500. Valid
// values are 1300 to 1700. For example, on a 180 degree servo, this would be
// 90 degrees. On continuous rotation servo, this is the 'neutral' position_sp
// where the motor does not turn. You must write to the position_sp attribute for
// changes to this attribute to take effect.
func (s *ServoMotor) MidPulseSp() int {
	return s.readInt("mid_pulse_sp")
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the miniumum (counter-clockwise) position_sp. Default value
// is 600. Valid values are 300 to 700. You must write to the position_sp
// attribute for changes to this attribute to take effect.
func (s *ServoMotor) SetMinPulseSp(x int) {
	s.writeInt("min_pulse_sp", x)
}

// Used to set the pulse size in milliseconds for the signal that tells the
// servo to drive to the miniumum (counter-clockwise) position_sp. Default value
// is 600. Valid values are 300 to 700. You must write to the position_sp
// attribute for changes to this attribute to take effect.
func (s *ServoMotor) MinPulseSp() int {
	return s.readInt("min_pulse_sp")
}

// Sets the polarity of the servo. Valid values are `normal` and `inversed`.
// Setting the value to `inversed` will cause the position_sp value to be
// inversed. i.e `-100` will correspond to `max_pulse_sp`, and `100` will
// correspond to `min_pulse_sp`.
func (s *ServoMotor) SetPolarity(x string) {
	s.writeString("polarity", x)
}

// Sets the polarity of the servo. Valid values are `normal` and `inversed`.
// Setting the value to `inversed` will cause the position_sp value to be
// inversed. i.e `-100` will correspond to `max_pulse_sp`, and `100` will
// correspond to `min_pulse_sp`.
func (s *ServoMotor) Polarity() string {
	return s.readString("polarity")
}

// Returns the name of the port that this motor is connected to.
func (s *ServoMotor) Address() string {
	return s.readString("address")
}

// Reading returns the current position_sp of the servo. Writing instructs the
// servo to move to the specified position_sp. Units are percent. Valid values
// are -100 to 100 (-100% to 100%) where `-100` corresponds to `min_pulse_sp`,
// `0` corresponds to `mid_pulse_sp` and `100` corresponds to `max_pulse_sp`.
func (s *ServoMotor) SetPositionSp(x int) {
	s.writeInt("position_sp", x)
}

// Reading returns the current position_sp of the servo. Writing instructs the
// servo to move to the specified position_sp. Units are percent. Valid values
// are -100 to 100 (-100% to 100%) where `-100` corresponds to `min_pulse_sp`,
// `0` corresponds to `mid_pulse_sp` and `100` corresponds to `max_pulse_sp`.
func (s *ServoMotor) PositionSp() int {
	return s.readInt("position_sp")
}

// Sets the rate_sp at which the servo travels from 0 to 100.0% (half of the full
// range of the servo). Units are in milliseconds. Example: Setting the rate_sp
// to 1000 means that it will take a 180 degree servo 2 second to move from 0
// to 180 degrees. Note: Some servo controllers may not support this in which
// case reading and writing will fail with `-EOPNOTSUPP`. In continuous rotation
// servos, this value will affect the rate_sp at which the speed ramps up or down.
func (s *ServoMotor) SetRateSp(x int) {
	s.writeInt("rate_sp", x)
}

// Sets the rate_sp at which the servo travels from 0 to 100.0% (half of the full
// range of the servo). Units are in milliseconds. Example: Setting the rate_sp
// to 1000 means that it will take a 180 degree servo 2 second to move from 0
// to 180 degrees. Note: Some servo controllers may not support this in which
// case reading and writing will fail with `-EOPNOTSUPP`. In continuous rotation
// servos, this value will affect the rate_sp at which the speed ramps up or down.
func (s *ServoMotor) RateSp() int {
	return s.readInt("rate_sp")
}

// Returns a list of flags indicating the state of the servo.
// Possible values are:
// * `running`: Indicates that the motor is powered.
func (s *ServoMotor) State() []string {
	return s.readStringArray("state")
}
