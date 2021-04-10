package exec

import (
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

type Exec struct {
	Command string
}

func NewExec(cmd string) *Exec {
	e := &Exec{
		Command: viper.GetString(cm.EXEC + "." + cmd),
	}
	return e
}

func (e *Exec) ExecVpn() {
	cmd := exec.Command(e.Command)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
