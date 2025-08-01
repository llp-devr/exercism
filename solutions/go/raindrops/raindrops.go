package raindrops

import (
    "strconv"
)
func Convert(number int) string {
	var buffer string
    if number % 3 == 0 {
        buffer = buffer + "Pling"
    }
	if number % 5 == 0 {
        buffer = buffer + "Plang"
    }
    if number % 7 == 0 {
        buffer = buffer + "Plong"
    }

    if buffer == "" {
        buffer = strconv.Itoa(number)
    }

    return buffer
    
}
