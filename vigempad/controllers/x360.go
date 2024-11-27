package controllers

import (
	"syscall"
	"unsafe"

	"github.com/CypherpunkSamurai/go-vigempad/vigempad"
)

// Xbox 360 Gamepad
type X360Gamepad struct {
	handle       uintptr
	clientHandle *vigempad.VigemClient
	connected    bool
	state        XUSB_REPORT
	model        GamepadModel
}

// Xbox 360 State Report
//
// Source Reference: https://github.com/nefarius/ViGEmClient/blob/b66d02d57e32cc8595369c53418b843e958649b4/include/ViGEm/Common.h#L69
type XUSB_REPORT struct {
	// uint16 representation of buttons
	wButtons uint16

	// uint8 range from 0-255
	TriggerL uint8
	// uint8 range from 0-255
	TriggerR uint8

	// int16 value range from -32768 to 32767
	ThumbLX int16
	// int16 value range from -32768 to 32767
	ThumbLY int16

	// int16 value range from -32768 to 32767
	ThumbRX int16
	// int16 value range from -32768 to 32767
	ThumbRY int16
}

// Xbox 360 Buttons
//
// Button constants for the Xbox 360 controller
// Bit Switch is used to set individual buttons
//
// https://www.geeksforgeeks.org/set-clear-and-toggle-a-given-bit-of-a-number-in-c/
//
// https://en.wikipedia.org/wiki/Bit_field
type Xbox360Buttons uint16

// https://www.rapidtables.com/convert/number/hex-to-binary.html?x=8000
const (

	// represented in 0000000000000001 binary
	Xbox360Buttons_UP Xbox360Buttons = 0x0001

	// represented in 0000000000000010 binary
	Xbox360Buttons_DOWN Xbox360Buttons = 0x0002

	// represented in 0000000000000100 binary
	Xbox360Buttons_LEFT Xbox360Buttons = 0x0004

	// represented in 0000000000001000 binary
	Xbox360Buttons_RIGHT Xbox360Buttons = 0x0008

	// represented in 0000000000010000 binary
	Xbox360Buttons_START Xbox360Buttons = 0x0010

	// represented in 0000000000100000 binary
	Xbox360Buttons_BACK Xbox360Buttons = 0x0020

	// represented in 0000000001000000 binary
	Xbox360Buttons_LEFT_THUMB Xbox360Buttons = 0x0040

	// represented in 0000000010000000 binary
	Xbox360Buttons_RIGHT_THUMB Xbox360Buttons = 0x0080

	// represented in 0000000100000000 binary
	Xbox360Buttons_LEFT_SHOULDER Xbox360Buttons = 0x0100

	// represented in 0000001000000000 binary
	Xbox360Buttons_RIGHT_SHOULDER Xbox360Buttons = 0x0200

	// represented in 0000010000000000 binary
	Xbox360Buttons_GUIDE Xbox360Buttons = 0x0400

	// represented in 0001000000000000 binary
	Xbox360Buttons_A Xbox360Buttons = 0x1000

	// represented in 0010000000000000 binary
	Xbox360Buttons_B Xbox360Buttons = 0x2000

	// represented in 0100000000000000 binary
	Xbox360Buttons_X Xbox360Buttons = 0x4000

	// represented in 1000000000000000 binary
	Xbox360Buttons_Y Xbox360Buttons = 0x8000
)

// Notification Update
type XUSB_NOTIFICATION func(
	Client uintptr,
	Target uintptr,
	largeMotor uint8,
	smallMotor uint8,
	ledNumber uint8,
	userData uintptr,
) uintptr

// Create a New Xbox 360 Gamepad
func NewX360Gamepad(client *vigempad.VigemClient) (*X360Gamepad, error) {
	handle, _, err := vigempad.TargetX360Alloc.Call()
	if vigempad.CheckSyscallError(err) {
		return nil, err
	}
	ret, _, err := vigempad.TargetAdd.Call(client.GetHandle(), handle)
	if vigempad.CheckSyscallError(err) {
		return nil, err
	}
	if err = vigempad.CheckVigemError(ret); err != nil {
		return nil, err
	}

	return &X360Gamepad{
		handle:       handle,
		clientHandle: client,
		connected:    true,
		model:        GamepadModel_X360Wired,
	}, nil
}

// Disconnect Xbox 360 Gamepad from the emulator
func (g *X360Gamepad) Disconnect() error {
	_, _, err := vigempad.TargetRemove.Call(g.GetHandle())
	if vigempad.CheckSyscallError(err) {
		return err
	}
	_, _, err = vigempad.TargetFree.Call(g.GetHandle())
	if vigempad.CheckSyscallError(err) {
		return err
	}
	g.connected = false
	return nil
}

// Get the ViGEm Client
func (g *X360Gamepad) GetClient() *vigempad.VigemClient {
	return g.clientHandle
}

// Get Gamepad Target Handle
func (g *X360Gamepad) GetHandle() uintptr {
	return g.handle
}

// Update Gamepad State
func (g *X360Gamepad) Update() error {
	ret, _, err := vigempad.TargetX360Update.Call(g.clientHandle.GetHandle(), g.handle, uintptr(unsafe.Pointer(&g.state)))
	if vigempad.CheckSyscallError(err) {
		return err
	}
	if err := vigempad.CheckVigemError(ret); err != nil {
		return err
	}
	return nil
}

// Press a Button on the Xbox 360 Gamepad
func (g *X360Gamepad) PressButton(button Xbox360Buttons) {
	g.state.wButtons = g.state.wButtons | uint16(button)
}

// Release a Button on Xbox 360 Gamepad
func (g *X360Gamepad) ReleaseButton(button Xbox360Buttons) {
	g.state.wButtons = g.state.wButtons & ^uint16(button)
}

// Set Right Thumbstick Position
//
// (value between -1 and 1, 0 is center)
func (g *X360Gamepad) SetThumbStickRight(x float32, y float32) {
	g.state.ThumbRX = int16(x * 32767)
	g.state.ThumbRY = int16(y * 32767)
}

// Set Left Thumbstick Position
//
// (value between -1 and 1, 0 is center)
func (g *X360Gamepad) SetThumbStickLeft(x float32, y float32) {
	g.state.ThumbLX = int16(x * 32767)
	g.state.ThumbLY = int16(y * 32767)
}

// Set Left Trigger Value
//
// (value between 0-255, 127 is 50% pressed)
func (g *X360Gamepad) SetTriggerLeft(value uint8) {
	g.state.TriggerL = value
}

// Set Right Trigger Value
//
// (value between 0-255, 127 is 50% pressed)
func (g *X360Gamepad) SetTriggerRight(value uint8) {
	g.state.TriggerR = value
}

// Register Update Callback
func (g *X360Gamepad) RegisterUpdateCallback(callback XUSB_NOTIFICATION) error {
	ret, _, err := vigempad.TargetX360RegisterNotification.Call(g.clientHandle.GetHandle(), g.handle, syscall.NewCallback(callback))
	if vigempad.CheckSyscallError(err) {
		return err
	}
	if err := vigempad.CheckVigemError(ret); err != nil {
		return err
	}
	return nil
}

// Unregister Update Callback
func (g *X360Gamepad) UnregisterUpdateCallback() error {
	_, _, err := vigempad.TargetX360UnregisterNotification.Call(g.clientHandle.GetHandle(), g.handle)
	if err != nil {
		return err
	}
	return nil
}
