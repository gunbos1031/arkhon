package cli

import (
	"flag"
	"os"
	"github.com/gunbos1031/arkhon/rest"
	"fmt"
)

func usage() {
	fmt.Println("Welcome to Arkhon Coin\n\n")
	fmt.Println("Please use the following flags:\n\n")
	fmt.Println("-port: 				Set the port of the server")
	fmt.Println("-mode:					Choose between 'web' and 'rest'")
	os.Exit(1)
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}
	
	port := flag.Int("port", 80, "Set port of the server")
	mode := flag.String("mode", "rest", "mode of the web we access")
	
	flag.Parse()
	
	switch *mode {
		case "rest":
		rest.Start(*port)
		case "web":
		default:
		usage()
	}
}

// go run main.go -mode= -port=