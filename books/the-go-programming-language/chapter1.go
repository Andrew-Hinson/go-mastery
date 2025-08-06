package books

import (
	"fmt"
	"os"
	"strings"
)

// Excercise4 (dup3) prints the count and text of lines that
// appear more than once in the named input files.

// Modify this function to print the names of all files in which each duplicated line occurs.
func Excercise4() {
	fileNameToPrint := ""
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		fileNameToPrint = filename
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n%s", n, line, fileNameToPrint)
		}
	}
}
