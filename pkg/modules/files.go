package modules

import (
	"encoding/json"
	"os"
)

type FileSystem struct {
	Directories []Directory
}

type Directory struct {
	Path string
	Mode string
}

type Files struct {
	fileSystem FileSystem
}

func (f *Files) PaseConfig() bool {
	content, err := os.ReadFile("test-data/foo.bar.Files")
	if err != nil {
		return false
	}
	println(string(content))
	err = json.Unmarshal(content, &f.fileSystem)
	return err == nil
}

func (f Files) Run(dryRun bool) string {
	for _, v := range f.fileSystem.Directories {
		println("mkdir " + v.Path)
		os.MkdirAll(v.Path, 0775)
	}

	return "xx"
}
