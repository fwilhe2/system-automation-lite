package modules

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Packages struct {
	packages []string
}

func (p *Packages) PaseConfig() bool {
	content, err := os.ReadFile("test-data/foo.bar.Packages")
	if err != nil {
		return false
	}
	println(string(content))
	p.packages = strings.Split(string(content), "\n")
	return true
}

func (p Packages) Run(dryRun bool) string {
	for _, pkg := range p.packages {
		fmt.Printf("Install package %s\n", pkg)
		out, err := exec.Command("sudo", "dnf", "-y", "install", pkg).CombinedOutput()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}

	return "xx"
}
