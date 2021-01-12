// Just for fun - Loading Screen Emulation for 0.001 seconds and then pause 2 seconds and refresh the Linux console
package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	// loading screen
	start := time.Now()
	for time.Since(start).Seconds() <= 0.001 { //for acts like while
		fmt.Println("Loading... Please Wait!")
	}

	// Sleep for 2 seconds
	time.Sleep(2 * time.Second)

	// Clear Screen
	cmd := exec.Command("clear") // command to console
	cmd.Stdout = os.Stdout       // command sent to console
	cmd.Run()                    // execute command
}
