package anagrams

import "fmt"

// Printer is used to print
type Printer interface {
	PrintAnagram(a []string)
	Print(a string)
}

// LoggerPrinter prints in logs
type LoggerPrinter struct {
	counter int
}

// PrintAnagram prints string array
func (cp *LoggerPrinter) PrintAnagram(an []string) {
	fmt.Println(an)
}

// Print prints string
func (cp *LoggerPrinter) Print(a string) {
	fmt.Println(a)
}
