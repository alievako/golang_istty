package main

import (
	"fmt"
	"log"
	"syscall"
)

func istty(fd int) (bool, error) {
	var stin syscall.Stat_t
	err := syscall.Fstat(fd, &stin)
	return stin.Mode&syscall.S_IFMT == syscall.S_IFCHR, err
}

func main() {
	for i := 0; i < 3; i++ {
		r, err := istty(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%t\n", r)
	}
}
