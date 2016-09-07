package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/naveego/api/pipeline/cli"
)

var (
	verbose = flag.Bool("v", false, "Verbose")
)

func main() {
	flag.Parse()

	if len(os.Args) < 4 {
		printUsage()
		return
	}

	pubOrSub := os.Args[1]
	name := os.Args[2]
	importPath := os.Args[3]

	var pkg cli.Package

	switch pubOrSub {
	case "publisher":
		pkg = cli.NewPublisherPackage(name, importPath)
	case "subscriber":
		pkg = cli.NewSubscriberPackage(name, importPath)
	default:
		printUsage()
		return
	}

	_, output, err := cli.BuildPackage(pkg)
	if err != nil {
		fmt.Printf("Could not create connector package: %v\n", err)
		fmt.Printf("Output: %s\n", output)
		os.Exit(1)
	}

	fmt.Println("Package create successfully")
}

func printUsage() {
	fmt.Println("Usage")
}
