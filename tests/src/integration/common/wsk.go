package common

import (
	"os"
	"os/exec"
)

const cmd = "wsk"
const arg = "-i"

type Wsk struct {
	Path string
	Arg []string
	Dir string
}

func NewWsk() *Wsk {
	return NewWskWithPath(os.Getenv("GOPATH") + "/src/github.com/openwhisk/openwhisk-cli/")
}

func NewWskWithPath(path string) *Wsk {
	var dep Wsk
	dep.Path = cmd
	dep.Arg = []string{arg}
	dep.Dir = path
	return &dep
}

func (wsk *Wsk)RunCommand(s ...string) ([]byte, error) {
	cs := wsk.Arg
	cs = append(cs, s...)
	command := exec.Command(wsk.Path, cs...)
	command.Dir = wsk.Dir
	return command.CombinedOutput()
}

func (wsk *Wsk)ListNamespaces() ([]byte, error) {
	return wsk.RunCommand("namespace", "list")
}