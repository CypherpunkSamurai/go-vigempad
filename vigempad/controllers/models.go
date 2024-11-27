package controllers

// List of Gamepad Models
type GamepadModel uint8

const (
	// Xbox 360 Controller (has an id of 0)
	GamepadModel_X360Wired GamepadModel = 0

	// DualShock 4 Controller (has an id of 2)
	GamepadModel_DS4Wired GamepadModel = 2
)
