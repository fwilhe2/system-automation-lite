package main

import "github.com/fwilhe2/system-automation-lite/pkg/modules"

type AutomationModule interface {
	PaseConfig() bool
	Run(dryRun bool) string
}

func main() {
	p := modules.Packages{}
	if p.PaseConfig() {
		p.Run(false)
	}

	f := modules.Files{}
	if f.PaseConfig() {
		f.Run(false)
	}
}
