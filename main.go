package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/Shyp/bump_version/lib"
)

const VERSION = "1.2"

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: bump_version [--version=<version>] [<major|minor|patch>] <filename>\n")
	flag.PrintDefaults()
}

// runCommand execs the given command and exits if it fails.
func runCommand(binary string, args ...string) {
	out, err := exec.Command(binary, args...).CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when running command: %s.\nOutput was:\n%s", err.Error(), string(out))
		os.Exit(2)
	}
}

var vsn = flag.String("version", "", "Set this version in the file (don't increment whatever version is present)")

func main() {
	flag.Usage = usage
	flag.Parse()
	args := flag.Args()
	var filename string
	var version *bump_version.Version
	if *vsn != "" {
		// no "minor"
		if len(args) != 1 {
			flag.Usage()
			return
		}
		var err error
		version, err = bump_version.Parse(*vsn)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(2)
		}
		filename = args[0]
		setErr := bump_version.SetInFile(version, filename)
		if setErr != nil {
			os.Stderr.WriteString(setErr.Error() + "\n")
			os.Exit(2)
		}
	} else {
		if len(args) != 2 {
			flag.Usage()
			return
		}
		versionTypeStr := args[0]
		filename = args[1]

		var err error
		version, err = bump_version.BumpInFile(bump_version.VersionType(versionTypeStr), filename)
		if err != nil {
			os.Stderr.WriteString(err.Error() + "\n")
			os.Exit(2)
		}
	}
	runCommand("git", "add", filename)
	runCommand("git", "commit", "-m", version.String())
	runCommand("git", "tag", version.String(), "--annotate", "--message", version.String())
	os.Stdout.WriteString(version.String() + "\n")
}
