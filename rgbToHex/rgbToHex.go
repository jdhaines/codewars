// The rgb function is incomplete. Complete it so that passing in RGB decimal values
// will result in a hexadecimal representation being returned. Valid decimal values for
// RGB are 0 - 255. Any values that fall out of that range must be rounded to the closest
// valid value.
//
// Note: Your answer should always be 6 characters long, the shorthand with 3 will not work here.
package kata_test

import (
	"fmt"
	"strconv"
)

// validateDecimalValue takes an integer and returns a valid 0-255 integer
func validateDecimalValue(input int) int {
	if input > 255 {
		return 255
	}
	if input < 0 {
		return 0
	}
	return input
}

// getDigitInHex takes an integer and returns the hex value as a string.
func getDigitInHex(input int) string {
	if input < 10 {
		return strconv.Itoa(input)
	}
	switch input {
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	default:
		return "This should never happen."
	}
}

// convertDecimalToHex takes an integer and returns the two-digit hex value as a string
func convertDecimalToHex(input int) string {
	validInput := validateDecimalValue(input)
	quotient := validInput / 16
	remainder := validInput % 16
	return fmt.Sprintf("%s%s", getDigitInHex(quotient), getDigitInHex(remainder))
}

// RGB takes three RBG integers and returns a six-character hex string for CSS
func RGB(r, g, b int) string {
	return fmt.Sprintf("%s%s%s", convertDecimalToHex(r), convertDecimalToHex(g), convertDecimalToHex(b))
}
