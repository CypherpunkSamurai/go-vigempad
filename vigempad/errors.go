package vigempad

import (
	"errors"
	"syscall"

	"golang.org/x/sys/windows"
)

// ViGEm Bus Error codes
//
// See Reference:
// https://github.com/fsadannn/pyvjoystick/blob/main/pyvjoystick/vigem/constants.py
type VigemError int

const (
	VigemErrorOk                        VigemError = 0x0        // Returned On Void
	VigemErrorNone                      VigemError = 0x20000000 // Returned when Success
	VigemErrorBusNotFound               VigemError = 0xE0000001 // Bus Not Found
	VigemErrorNoFreeSlot                VigemError = 0xE0000002 // No Free Slot
	VigemErrorInvalidTarget             VigemError = 0xE0000003 // Invalid Target
	VigemErrorRemovalFailed             VigemError = 0xE0000004 // Removal Failed
	VigemErrorAlreadyConnected          VigemError = 0xE0000005 // Already Connected
	VigemErrorTargetUninitialized       VigemError = 0xE0000006 // Target Uninitialized
	VigemErrorTargetNotPluggedIn        VigemError = 0xE0000007 // Target Not Plugged In
	VigemErrorBusVersionMismatch        VigemError = 0xE0000008 // Bus Version Mismatch
	VigemErrorBusAccessFailed           VigemError = 0xE0000009 // Bus Access Failed
	VigemErrorCallbackAlreadyRegistered VigemError = 0xE0000010 // Callback Already Registered
	VigemErrorCallbackNotFound          VigemError = 0xE0000011 // Callback Not Found
	VigemErrorBusAlreadyConnected       VigemError = 0xE0000012 // Bus Already Connected
	VigemErrorBusInvalidHandle          VigemError = 0xE0000013 // Bus Invalid Handle
	VigemErrorXusbUserindexOutOfRange   VigemError = 0xE0000014 // Xusb Userindex Out Of Range
	VigemErrorInvalidParameter          VigemError = 0xE0000015 // Invalid Parameter
	VigemErrorNotSupported              VigemError = 0xE0000016 // Not Supported
)

// Check if the Error is a Syscall Error
func CheckSyscallError(err error) bool {
	if err != nil && err != syscall.Errno(windows.ERROR_SUCCESS) {
		return true
	}
	return false
}

// Check if the return value is a ViGEm Error
func CheckVigemError(ret uintptr) error {
	// log checking
	// log.Printf("[checkVigemError] checking return value \"0x%x\" for error.\n", ret)
	// check return value
	switch VigemError(ret) {
	case VigemErrorOk:
		return nil
	case VigemErrorNone:
		return nil
	case VigemErrorBusNotFound:
		return errors.New("bus not found")
	case VigemErrorNoFreeSlot:
		return errors.New("no free slot")
	case VigemErrorInvalidTarget:
		return errors.New("invalid target")
	case VigemErrorRemovalFailed:
		return errors.New("removal failed")
	case VigemErrorAlreadyConnected:
		return errors.New("already connected")
	case VigemErrorTargetUninitialized:
		return errors.New("target uninitialized")
	case VigemErrorTargetNotPluggedIn:
		return errors.New("target not plugged in")
	case VigemErrorBusVersionMismatch:
		return errors.New("bus version mismatch")
	case VigemErrorBusAccessFailed:
		return errors.New("bus access failed")
	case VigemErrorCallbackAlreadyRegistered:
		return errors.New("callback already registered")
	case VigemErrorCallbackNotFound:
		return errors.New("callback not found")
	case VigemErrorBusAlreadyConnected:
		return errors.New("bus already connected")
	case VigemErrorBusInvalidHandle:
		return errors.New("bus invalid handle")
	case VigemErrorXusbUserindexOutOfRange:
		return errors.New("xusb user index out of range")
	case VigemErrorInvalidParameter:
		return errors.New("invalid parameter")
	case VigemErrorNotSupported:
		return errors.New("operation not supported")
	default:
		return nil
	}
}
