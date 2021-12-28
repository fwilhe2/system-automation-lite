package main

import (
	"encoding/json"
	"fmt"
	// "io/fs"
	"os"
	// "os/exec"
	// "strconv"
	// "strings"
)

type Playbook struct {
	Meta  Meta
	Tasks map[string]Task
}

type Task interface {
	name() string
}

type TaskDescription interface {
	listTasks() []Task
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

func (f FileSystem) name() string {
	return f.Directories[0].Path
}

type Directory struct {
	Path string
	Mode string
}

type Packages struct {
	DebianFamily []string
	RedhatFamily []string
}

func (f Packages) name() string {
	return f.DebianFamily[0]
}

func f(p Task) {
	fmt.Printf("f: %+v\n", p)

	println(p.name())

}

func g(p Task) {
	fmt.Printf("g: %+v\n", p)

	println(p.name())

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

    mapping := map[string]func(Task){
        "FileSystem": f,
        "Packages": g,
    }

	for k, v := range playbook.Tasks {
		println(k)
		println(v)

		myFunc := mapping[k]

		myFunc(v)
	}

	// for _, d := range playbook.Tasks.FileSystem.Directories {
	// 	mode, err := strconv.Atoi(d.Mode)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("mkdir %s %s\n", d.Path, d.Mode)
	// 	err = os.MkdirAll(d.Path, fs.FileMode(mode))
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// osReleaseFile, err := os.ReadFile("/etc/os-release")
	// if err != nil {
	// 	panic(err)
	// }
	// osReleaseString := string(osReleaseFile)
	// rhLike := false
	// debianLike := false
	// for _, v := range strings.Split(osReleaseString, "\n") {
	// 	entry := strings.Split(v, "=")
	// 	if entry[0] == "ID" {
	// 		if entry[1] == "debian" || entry[1] == "ubuntu" {
	// 			fmt.Println("OS is debian-like")
	// 			debianLike = true
	// 		}

	// 		if entry[1] == "fedora" || entry[1] == "centos" {
	// 			fmt.Println("OS is rh-like")
	// 			rhLike = true
	// 		}
	// 	}
	// }

	// if debianLike {
	// 	for _, v := range playbook.Tasks.Packages.DebianFamily {
	// 		fmt.Printf("Install package %s\n", v)
	// 		out, err := exec.Command("sudo", "apt-get", "-y", "install", v).CombinedOutput()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Println(string(out))
	// 	}
	// }

	// if rhLike {
	// 	for _, v := range playbook.Tasks.Packages.RedhatFamily {
	// 		fmt.Printf("Install package %s\n", v)
	// 		out, err := exec.Command("sudo", "dnf", "-y", "install", v).CombinedOutput()
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		fmt.Println(string(out))
	// 	}
	// }

}
