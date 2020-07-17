/*
Some helper functions used in various parts of the system.
*/
package utilities

import (
	"bytes"
	"fmt"
)

// Returns true if the given array of byte arrays contains lookFor.
func ByteArrayContains(array [][]byte, lookFor []byte) bool {
	for _, item := range array {
		if bytes.Equal(item, lookFor) {
			return true
		}
	}
	return false
}

// Used to find a cycle, lookFor must be at len-1 or len-2 of the array.
func CycleContains(array [][]byte, lookFor []byte) bool {
	return (len(array) > 0 && bytes.Equal(array[len(array)-1], lookFor)) || (len(array) > 1 && bytes.Equal(array[len(array)-2], lookFor))
}

// Compares two byte arrays for sorting purposes.
// Return: true if a should come first in ascending order, false otherwise.
func ByteArrayComparator(a, b []byte) bool {
	var byte1, byte2 int
	result := false
	for k := 0; k < len(a) && k < len(b); k++ {
		byte1 = int(a[k]) & 0xff
		byte2 = int(b[k]) & 0xff
		result = byte1 < byte2
		if byte1 != byte2 {
			break
		}
	}
	return result
}

// Make a clean copy of the <source> byte array with the given length to reduce memory footprint.
func ByteArrayCopy(source []byte, length int) []byte {
	result := make([]byte, length)
	copy(result, source)
	return result
}

func ByteArrayToString(a []byte) string {
	if a == nil {
		return "(null)"
	} else if len(a) == 0 {
		return "(empty)"
	} else if len(a) <= 4 {
		return fmt.Sprintf("%x", a)
	} else {
		return fmt.Sprintf("%x...%x", a[0:2], a[len(a)-2:])
	}
}
