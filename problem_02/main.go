package main

import (
	"fmt"
	"os/exec"
	"time"
)

/*
	===========================================================================================
	Proposed Solution
	===========================================================================================
	Procedure:
		So the linux command to rename a file called original.txt to changed.txt is as follows:
			"mv original.txt changed.txt"
	Here is a simple go code that runs as follows:
		1) Delete changed.txt if exists from previous runs.
		2) Create the "original.txt" file,
		3) Wait 3 seconds to see the file clearly in the file explorer.
		4) Proceed to run the "mv" linux command to rename it to changed.txt
*/

func main() {
	oldName := "original.txt"
	newName := "changed.txt"

	// Delete changed.txt if exists from previous runs.
	cmd := exec.Command("rm", "-f", newName)
	_ = cmd.Run()

	// Create the "original.txt" file
	cmd = exec.Command("touch", oldName)
	_ = cmd.Run()

	// Wait 3 seconds to see the file clearly in the file explorer.
	time.Sleep(3 * time.Second)

	// Run the "mv" linux command to rename it to changed.txt
	cmd = exec.Command("mv", oldName, newName)
	_ = cmd.Run()

	fmt.Println("File renamed successfully using Linux command!")
}
