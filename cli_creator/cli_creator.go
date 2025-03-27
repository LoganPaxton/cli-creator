package cli_creator

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var GLOBAL_BUFF_STORAGE string
var GLOBAL_ARG_BUFFER []string
var GLOBAL_COMMAND_LIST = make(map[string]func()) // Map to store commands and their handlers

// Read a line from input
func ReadLine() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	GLOBAL_BUFF_STORAGE = text
	return GLOBAL_BUFF_STORAGE
}

// Split input into arguments
func FetchArgs() []string {
	GLOBAL_ARG_BUFFER = []string{} // Reset arguments

	if GLOBAL_BUFF_STORAGE == "" {
		return GLOBAL_ARG_BUFFER
	}

	GLOBAL_ARG_BUFFER = strings.Fields(GLOBAL_BUFF_STORAGE) // Splits on spaces automatically
	return GLOBAL_ARG_BUFFER
}

// Register a new command
func RegisterCommand(command string, action func()) {
	GLOBAL_COMMAND_LIST[command] = action
}

// Parse and execute commands
func ParseArgs() {
	if len(GLOBAL_ARG_BUFFER) == 0 {
		return
	}

	cmd := GLOBAL_ARG_BUFFER[0] // First argument is the command

	if action, exists := GLOBAL_COMMAND_LIST[cmd]; exists {
		action() // Execute the associated function
	} else {
		fmt.Println("Unknown command:", cmd)
	}
}
