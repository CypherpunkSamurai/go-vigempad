package main

import (
	"fmt"
	"log"
	"time"

	"github.com/CypherpunkSamurai/go-vigempad/vigempad"
	"github.com/CypherpunkSamurai/go-vigempad/vigempad/controllers"
)

func main() {
	// Init the ViGem Client first
	fmt.Println("Initializing ViGem Client")

	// create a new ViGEm Client
	vigemClient, err := vigempad.NewVigemClient()
	if err != nil {
		log.Fatalln("Error initializing ViGem Client:", err)
	}
	fmt.Println("failed")

	// create a new Xbox 360 controller
	x360controller, err := controllers.NewX360Gamepad(vigemClient)
	if err != nil {
		log.Fatalln("Error initializing Xbox 360 controller:", err)
	}

	// wait 1 second before clicking button
	fmt.Println("Pressing Triggers in 1 second")
	time.Sleep(time.Second)

	// press joysticks test
	for i := 0.0; i < 1.0; i += 0.5 {
		x360controller.SetThumbStickLeft(float32(i), float32(i))
		x360controller.SetThumbStickRight(float32(i), float32(i))
		time.Sleep(100 * time.Millisecond)
	}

	// wait 1 second before releasing button
	fmt.Println("Disconnecting controller in 1 second")
	time.Sleep(time.Second)

	// disconnect the controller
	x360controller.Disconnect()

	// disconnect the client
	vigemClient.Disconnect()
}
