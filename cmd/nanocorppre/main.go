package main

import "github.com/attilasatan/nanocorp"

func main() {

	if err := nanocorp.Init(); err != nil {
		panic(err) //TODO: add some fancy error handling logic.
	}
}
