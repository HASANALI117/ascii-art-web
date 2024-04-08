package asciiArt

import "fmt"

func StartLineCalc(char rune) (int, error) {
	if char < ' ' || char > '~' {
		return -1, fmt.Errorf("ERROR: %c is not a valid ascii character\n", char)
	}
	return (int(char)-' ')*9 + 2, nil
}
