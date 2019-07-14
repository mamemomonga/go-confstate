package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/mamemomonga/go-confstate/examples/myconfstate"
	"github.com/sethvargo/go-password/password"
	"log"
)

func main() {
	// Initalize and load configs and states
	if err := myconfstate.Load(); err != nil {
		log.Fatal(err)
	}

	// Set a password for a user whose password is not set
	for _, name := range myconfstate.C().Users {
		if _, ok := myconfstate.S().Passwords[name]; !ok {
			pwd, err := password.Generate(16, 2, 1, false, false)
			if err != nil {
				log.Fatal(err)
			}
			myconfstate.S().Passwords[name] = pwd
		}
	}

	// Save StatesFile
	if err := myconfstate.Save(); err != nil {
		log.Fatal(err)
	}

	// Dump it
	spew.Dump(myconfstate.S())
}
