package main

import (
    "os"
	"log"

    "svc-dynamic-form-go/framework/commands"
)

func main() {

    if len(os.Args) < 2 {
		log.Fatal("Please provide a command.")
	}

	cmd := os.Args[1]

	switch cmd {
	case "serve":
		commands.Serve()
	case "bootstrap":
		commands.Bootstrap()
    case "make":
		commands.Make(os.Args[2:])
	case "migrate":
        commands.Migrate()
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}

}
