package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const usage = "Usage: ./debian-version-calculator <git checkout directory> [force prerelease]"

func main() {
	forcePreRelease := false

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

	if os.Args[1] == "--help" || os.Args[1] == "-h" {
		fmt.Println(usage)
		return
	}

	if len(os.Args) > 2 {
		forcePreRelease = strings.ToLower(os.Args[2]) == "true"
	}

	dir := os.Args[1]

	result := &upstream{}

	_, err := pkgVersionFromGit(dir, result, forcePreRelease)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", result.version)
	fmt.Printf("%v\n", result.tag)
	fmt.Printf("%v\n", result.commitIsh)
	fmt.Printf("%v\n", result.hasRelease)
	fmt.Printf("%v\n", result.isRelease)
}
