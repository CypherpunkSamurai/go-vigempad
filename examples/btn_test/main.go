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
	fmt.Println("waiting 1 second before clicking button A")
	time.Sleep(time.Second)

	// press button
	x360controller.PressButton(controllers.Xbox360Buttons_A)

	// wait 1 second before releasing button
	fmt.Println("waiting 1 second before releasing button A")
	time.Sleep(time.Second)

	// release button
	x360controller.ReleaseButton(controllers.Xbox360Buttons_A)

	// disconnect the controller
	x360controller.Disconnect()

	// disconnect the client
	vigemClient.Disconnect()
}
