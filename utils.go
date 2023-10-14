package main

import (
	"bufio"
	"os"
)

func ClearInputBuffer() {
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

