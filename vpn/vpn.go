package vpn

import (
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

type Vpn struct {
	Command string
}

func NewVpn(provider string) *Vpn {
	v := &Vpn{
		Command: viper.GetString(cm.VPN + "." + provider + "." + cm.CMD),
	}
	return v
}

func (v *Vpn) ExecVpn() {
	cmd := exec.Command(v.Command)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
