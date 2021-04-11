package profile

import (
	"fmt"
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/alknopfler/alkcli/connect"
	"github.com/alknopfler/alkcli/exec"
	"github.com/alknopfler/alkcli/tunnel"
	"github.com/alknopfler/alkcli/vpn"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Profile struct {
	ParentCMD map[string]map[string]string
}

func NewProfile(name string) *Profile {
	p := &Profile{
		ParentCMD: viper.GetStringMap(cm.PROFILE + "." + name),
	}
	fmt.Println(p.ParentCMD)
	return p
}

func (p *Profile) ExecProfile() {

	for i, val := range p.ParentCMD {
		switch p.ParentCMD[i] {
		case val["connection"]:
			connect.NewConnection(execution).ExecConnection()
			time.Sleep(2)
		case "tunnel":
			tunnel.NewTunnel(execution)
		case "vpn":
			vpn.NewVpn(execution)
			time.Sleep(10) //just to wait until finish connection
		case "exec":
			exec.NewExec(execution)
			time.Sleep(3) //just to wait until finish execution
		}
	}
}
