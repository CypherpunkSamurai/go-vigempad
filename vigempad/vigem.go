package vigempad

import (
	"errors"
	"os"
)

// ViGEm Bus Client
type VigemClient struct {
	handle uintptr
}

// Create a New ViGEm Bus Client
func NewVigemClient() (*VigemClient, error) {

	// check if ViGEmClient.dll is available
	if _, err := os.Stat("ViGEmClient.dll"); errors.Is(err, os.ErrNotExist) {
		return nil, errors.New("ViGEmClient.dll not found")
	}

	// allocate a new vigem client
	clientHandle, _, err := Alloc.Call()
	if CheckSyscallError(err) {
		return nil, err
	}

	// connect the vigem bus
	ret, _, err := Connect.Call(clientHandle)
	if CheckSyscallError(err) {
		return nil, err
	}
	if CheckVigemError(ret) != nil {
		return nil, err
	}

	return &VigemClient{
		handle: clientHandle,
	}, nil
}

// Disconnect a ViGEm Client
func (c *VigemClient) Disconnect() error {
	_, _, err := Disconnect.Call(c.handle)
	if CheckSyscallError(err) {
		return err
	}

	// Free ViGEm Client
	_, _, err = Free.Call(c.handle)
	if CheckSyscallError(err) {
		return err
	}

	return nil
}

// GetHandle returns the uintptr of the ViGEm Client
func (c *VigemClient) GetHandle() uintptr {
	return c.handle
}
