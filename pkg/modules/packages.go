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
	content, err := os.ReadFile("foo.bar.Packages")
	if err != nil {
		return false
	}
	println(string(content))
	p.packages = strings.Split(string(content), "\n")
	return true
}

func (p Packages) Run(dryRun bool) string {
	osReleaseFile, err := os.ReadFile("/etc/os-release")
	if err != nil {
		panic(err)
	}
	osReleaseString := string(osReleaseFile)
	rhLike := false
	debianLike := false
	for _, v := range strings.Split(osReleaseString, "\n") {
		entry := strings.Split(v, "=")
		if entry[0] == "ID" {
			if entry[1] == "debian" || entry[1] == "ubuntu" {
				fmt.Println("OS is debian-like")
				debianLike = true
			}

			if entry[1] == "fedora" || entry[1] == "centos" {
				fmt.Println("OS is rh-like")
				rhLike = true
			}
		}
	}

	if debianLike {
		for _, pkg := range p.packages {
			fmt.Printf("Install package %s\n", pkg)
			out, err := exec.Command("sudo", "apt-get", "-y", "install", pkg).CombinedOutput()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		}
	}

	if rhLike {
		for _, pkg := range p.packages {
			fmt.Printf("Install package %s\n", pkg)
			out, err := exec.Command("sudo", "dnf", "-y", "install", pkg).CombinedOutput()
			if err != nil {
				panic(err)
			}
			fmt.Println(string(out))
		}
	}

	return "xx"
}
