package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
)

func main() {
	for {
		// Retrieve virtual memory information
		virtualMemory, _ := mem.VirtualMemory()

		// Clear the terminal screen and move the cursor to the top-left corner
		fmt.Printf("\033[2J")
		fmt.Printf("\033[1;1H")

		// Display available memory and memory allocated for the program
		fmt.Printf("Available Memory: %d MiB\n", virtualMemory.Available>>20)
		fmt.Printf("Allocated for this program: %d MiB\n", len(global_arr)<<7)

		// Check if available memory is less than 256 MiB
		if virtualMemory.Available>>20 < 256 {
			fmt.Printf("Memory is not enough, allocated memory has been skipped")
			time.Sleep(time.Second * 5)
			continue
		}

		// Allocate memory and display information
		allocate_arr()
		time.Sleep(time.Second)
	}
}

// global_arr is a slice to keep track of allocated memory
var global_arr = []interface{}{}

// allocate_arr appends a new 128 MiB array to global_arr
func allocate_arr() {
	global_arr = append(global_arr, _128m_arr())
}

// _128m_arr creates and returns a 128 MiB byte array
func _128m_arr() interface{} {
	arr := [1 << 27]byte{} // 1 << 27 is equivalent to 128 MiB
	for i := 0; i < 1<<27; i++ {
		arr[i] = 0
	}
	return &arr
}
