package term

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// YesNo prompts user to enter Y (yes) or N (no) and returns boolean
// respectively. If input is invalid, error is returned.
func YesNo(prompt string, defaultValue bool) (bool, error) {
	reader := bufio.NewReader(os.Stdin)

	defaultAnswer := "y/N"
	if defaultValue {
		defaultAnswer = "Y/n"
	}

	fmt.Printf("%s [%s]: ", prompt, defaultAnswer)

	input, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}

	input = strings.TrimSpace(input)

	if input == "" {
		return defaultValue, nil
	}

	switch strings.ToLower(input) {
	case "y", "yes":
		return true, nil
	case "n", "no":
		return false, nil
	default:
		return false, fmt.Errorf("invalid input: %s", input)
	}
}
