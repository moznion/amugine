package main

import (
	"os"

	"github.com/moznion/amugine"
)

func main() {
	amugine.Run(os.Args[1:])
}
