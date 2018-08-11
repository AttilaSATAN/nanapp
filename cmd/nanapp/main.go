// Package main is just a simple CLI interface for `github.com/attilasatan/nanapp` package.
package main

import (
	"errors"
	"log"
	"os"

	"github.com/attilasatan/nanapp"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal(errors.New("data file didn't provided"))
	} else if err := nanapp.Init(os.Args[1]); err != nil {
		log.Fatal(err) //TODO: add some fancy error handling logic.
	}

}
