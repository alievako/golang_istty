package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"
)

func istty(fd int) (bool, error) {
	var stin syscall.Stat_t
	err := syscall.Fstat(fd, &stin)
	return stin.Mode&syscall.S_IFMT == syscall.S_IFCHR, err
}

func main() {

	arg, err := strconv.Atoi(os.Args[1])
	check(err)

	r, err := istty(arg)
	check(err)

	fmt.Printf("%t\n", r)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
