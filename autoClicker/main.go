package main

import (
	"fmt"
	"github.com/go-vgo/robotgo" // This library will help to click
	"github.com/robotn/gohook"  // This library will help to listen for the keyboard
	"time"
)

func main() {
	fmt.Println("Auto Clicker")
	s := hook.Start()          // Starting the hooker
	defer hook.End()           // Executing hook.End() at the end
	autoClickerActive := false // Variable for checking if auto clicker is active or not
	go func() {                // An infinite loop that runs concurrently and it will check whether if the auto clicker is active or not, if so, then it will click
		for {
			if autoClickerActive {
				robotgo.MouseClick("left", false) // "left" tells that I need it to left click. false tells that we need single click and not a double click
			}
			time.Sleep(time.Millisecond * 5) // Delay of 5ms for each iteration. Setting it lower than 5 may crash the system
		}
	}()
	fmt.Println("Press F6 to Enable/Disable the Auto Clicker")
	for { // A for loop that will listen for the F6 key
		select {
		case i := <-s:
			if i.Kind > hook.KeyDown && i.Kind < hook.KeyUp { // Checking for a key press
				if i.Rawcode == 117 { // Checking if the rawcode matches with the rawcode of the F6 key (117)
					autoClickerActive = !autoClickerActive // Inverting autoClickerActive's value
					if autoClickerActive {
						fmt.Println("Enabled")
					} else {
						fmt.Println("Disabled")
					}
				}
			}
		}
	}
}
