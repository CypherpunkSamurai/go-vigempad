package vigempad

// ViGEm Gamepad Interface
type VigemGampadInterface interface {
	GetHandle() uintptr
	GetClient() *VigemClient
	Disconnect() error
	Update() error
}
