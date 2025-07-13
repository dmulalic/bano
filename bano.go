package main

import (
	"flag"
	"fmt"
)

func InitFlags() {
	flag.Usage =func() {
	fmt.Println("Usage: bano [OPTIONS] [FILE]...")
	}
}

func main() {

InitFlags()
}
