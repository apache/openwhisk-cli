package common

import (
    "os"
    "os/exec"
    "strings"
)

const cmd = "./wsk"
const arg = "-i"

type Wsk struct {
    Path string
    Arg []string
    Dir string
    Wskprops *Wskprops
}

func NewWsk() *Wsk {
    return NewWskWithPath(os.Getenv("GOPATH") + "/src/github.com/openwhisk/openwhisk-cli/")
}

func NewWskWithPath(path string) *Wsk {
    var dep Wsk
    dep.Path = cmd
    dep.Arg = []string{arg}
    dep.Dir = path
    dep.Wskprops = GetWskprops()
    return &dep
}

func (wsk *Wsk)Exists() bool {
    _, err := os.Stat(wsk.Dir + wsk.Path);
    if err == nil {
        return true
    } else {
        return false
    }
}

func (wsk *Wsk)RunCommand(s ...string) ([]byte, error) {
    cs := wsk.Arg
    cs = append(cs, s...)
    command := exec.Command(wsk.Path, cs...)
    command.Dir = wsk.Dir
    return command.CombinedOutput()
}

func (wsk *Wsk)ListNamespaces() ([]byte, error) {
    return wsk.RunCommand("namespace", "list", "--apihost", wsk.Wskprops.APIHost,
        "--auth", wsk.Wskprops.AuthKey)
}

func (wsk *Wsk)Getfqn(name string) string {
    sep := "/"
    if (strings.HasPrefix(name, sep)) {
        return name
    } else {
        return sep + wsk.Wskprops.Namespace + sep + name
    }
}

func (wsk *Wsk)CreateAction(name string, artifact string) ([]byte, error) {
    return wsk.CreateUpdateAction(name, artifact, false)
}

func (wsk *Wsk)UpdateAction(name string, artifact string) ([]byte, error) {
    return wsk.CreateUpdateAction(name, artifact, true)
}

func (wsk *Wsk)CreateUpdateAction(name string, artifact string, update bool) ([]byte, error) {
    arg := []string{"action"}
    if update {
        arg = append(arg, "update")
    } else {
        arg = append(arg, "create")
    }
    arg = append(arg, wsk.Getfqn(name), artifact, "--apihost", wsk.Wskprops.APIHost,
        "--auth", wsk.Wskprops.AuthKey, "--apiversion", wsk.Wskprops.APIVersion)
    return wsk.RunCommand(arg...)
}

func (wsk *Wsk)DeleteAction(name string) ([]byte, error) {
    return wsk.RunCommand("action", "delete", wsk.Getfqn(name), "--apihost", wsk.Wskprops.APIHost,
        "--auth", wsk.Wskprops.AuthKey, "--apiversion", wsk.Wskprops.APIVersion)
}
