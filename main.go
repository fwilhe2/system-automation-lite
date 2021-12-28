package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"strconv"
)

type Playbook struct {
	Meta  Meta
	Tasks Tasks
}

type Meta struct {
	Name        string
	Description string
}

type Tasks struct {
	FileSystem FileSystem
	Packages   Packages
}

type FileSystem struct {
	Directories []Directory
}

type Directory struct {
	Path string
	Mode string
}

type Packages struct {
	DebianFamily []string
	RedhatFamily []string
}

func main() {
	fileByte, err := os.ReadFile("test-data/simple.json")
	if err != nil {
		panic(err)
	}
	var playbook = Playbook{}
	err = json.Unmarshal(fileByte, &playbook)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", playbook)

	for _, d := range playbook.Tasks.FileSystem.Directories {
		mode, err := strconv.Atoi(d.Mode)
		if err != nil {
			panic(err)
		}
		fmt.Printf("mkdir %s %s\n", d.Path, d.Mode)
		err = os.MkdirAll(d.Path, fs.FileMode(mode))
		if err != nil {
			panic(err)
		}
	}

	for _, v := range playbook.Tasks.Packages.DebianFamily {
		fmt.Printf("todo: install package %s\n", v)
	}
}
