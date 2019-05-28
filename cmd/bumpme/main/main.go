package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/rogozhka/bumpme/internal/semver"
)

const VERSION = "1.4.0"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bumpme [--version=<version>] [<major|minor|patch>] <filename>\n")
	flag.PrintDefaults()
}

// runCommand execs the given command and exits if it fails.
func runCommand(cmd string, args ...string) {
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		log.Fatalf("cannot run | %v | %v | %v", cmd, out, err)
	}
}

var vsn = flag.String("version", "", "Alter value for 'const version='x.y.z'")
var extract = flag.Bool("extract", false, "Extract version from main.go and print")

func main() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.Ldate)

	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	var filename string
	var version *semver.Version

	if *vsn != "" {
		// no "minor"
		if len(args) != 1 {
			flag.Usage()
			return
		}
		var err error
		version, err = semver.Parse(*vsn)
		if err != nil {
			log.Fatalf("cannot parse | %v", err)
		}

		filename = args[0]
		setErr := semver.SetInFile(version, filename)
		if setErr != nil {
			log.Fatalf("cannot SetInFile | %v", err)
		}
	} else if *extract == true {
		ver, err := semver.Extract(args[0])
		if err != nil {
			log.Fatalf("cannot Extract | %v", err)
		}
		os.Stdout.WriteString(ver)
	} else {
		if len(args) != 2 {
			flag.Usage()
			return
		}
		versionTypeStr := args[0]
		filename = args[1]

		var err error
		version, err = semver.BumpInFile(semver.VersionType(versionTypeStr), filename)
		if err != nil {
			log.Fatalf("cannot BumpInFile | %v", err)

		}
		os.Stdout.WriteString(version.String() + "\n")
	}
}
