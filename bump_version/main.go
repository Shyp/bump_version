package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Shyp/bump_version"
)

const VERSION = "4.1.1"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bump_version <major|minor|patch> <filename>\n")
}

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		usage()
		os.Exit(2)
	}
	versionTypeStr := args[0]
	filename := args[1]

	version, err := bump_version.BumpInFile(bump_version.VersionType(versionTypeStr), filename)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Fprintf(os.Stderr, "Bumped version to %s\n", version)
	}
	out, err := exec.Command("git", "tag", version.String()).CombinedOutput()
	if err != nil {
		log.Fatalf("Error when attempting to git tag: %s.\nOutput was:\n%s", err.Error(), string(out))
	}
	fmt.Fprintf(os.Stderr, "Tagged git version: %s. Commit your changes\n", version)
}
