package vigempad

import "syscall"

// ViGEmClient Dll
var dll = syscall.NewLazyDLL("ViGEmClient.dll")

var (

	// Allocate a ViGEm Bus Handle
	//
	// Call() (handle uintptr, void, error)
	Alloc = dll.NewProc("vigem_alloc")

	// Connect a ViGEm Bus Handle
	//
	// Call(handle uintptr) (ViGEmError, void, error)
	Connect = dll.NewProc("vigem_connect")

	// Disconnect Vigem Bus
	//
	// Call(handle uintptr) (void, void, error)
	Disconnect = dll.NewProc("vigem_disconnect")

	// Free ViGEm Bus Handle from Memory
	//
	// Call(handle uintptr) (void, void, void)
	Free = dll.NewProc("vigem_free")

	// Add a ViGEm Bus Target Gamepad
	//
	// Call(clientHandle uintptr, targetHandle uintptr) (ViGEmError, void, error)
	TargetAdd = dll.NewProc("vigem_target_add")

	// Returns Boolean Whether Target is Connected
	//
	// Call(handle uintptr) (bool, void, error)
	TargetIsAttached = dll.NewProc("vigem_target_is_attached")

	// Remove a Target from ViGEm Bus
	//
	// Call(clientHandle uintptr, targetHandle uintptr) (ViGEmError, void, error)
	TargetRemove = dll.NewProc("vigem_target_remove")

	// Free a target from ViGEm Bus
	//
	// Call(handle uintptr) (ViGEmError, void, error)
	TargetFree = dll.NewProc("vigem_target_free")
)

// Xbox 360 Gamepad Target
//
// Xbox 360 Gamepad has a ViGEm ID of 0
var (
	// Allocate a Xbox 360 Gamepad Target
	//
	// Call() (handle uintptr, void, error)
	TargetX360Alloc = dll.NewProc("vigem_target_x360_alloc")

	// Update Xbox 360 Gamepad Target State
	//
	// Call(clientHandle uintptr, targetHandle uintptr, state *Xbox360State) (ViGEmError, void, error)
	//
	// Where unsafe uintptr *Xbox360State{} struct contains {wButtons, bLeftTrigger, bRightTrigger, sThumbLX, sThumbLY, sThumbRX, sThumbRY}
	TargetX360Update = dll.NewProc("vigem_target_x360_update")

	// Register a Notification Callback for Xbox 360 Gamepad Target
	//
	// Call(clientHandle uintptr, targetHandle uintptr, notification Callback)
	TargetX360RegisterNotification = dll.NewProc("vigem_target_x360_register_notification")

	// Unregister Notification Callback for Xbox 360 Gamepad Target
	//
	// Call(targetHandle uintptr) (void, error)
	TargetX360UnregisterNotification = dll.NewProc("vigem_target_x360_unregister_notification")
)

// DualShock 4 Gamepad Target
//
// DualShock 4 Gamepad has a ViGEm ID of 2
var (
	// Alocate a DualShock 4 Gamepad Target
	//
	// Call() (handle uintptr, void, error)
	TargetDs4Alloc = dll.NewProc("vigem_target_ds4_alloc")

	// Update DualShock 4 Gamepad Target State
	//
	// Call(clientHandle uintptr, targetHandle uintptr, state *Ds4State) (ViGEmError, void, error)
	//
	// Where unsafe uintptr *Ds4State{} struct contains {wButtons, bTriggerL, bTriggerR, sThumbLX, sThumbLY, sThumbRX, sThumbRY, wButtonsSpecial}
	TargetDs4Update = dll.NewProc("vigem_target_ds4_update")
)
