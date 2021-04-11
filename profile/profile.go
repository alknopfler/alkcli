package profile

import (
	"fmt"
	cm "github.com/alknopfler/alkcli/configMgmt"
	conn "github.com/alknopfler/alkcli/connect"
	"github.com/alknopfler/alkcli/exec"
	"github.com/alknopfler/alkcli/tunnel"
	"github.com/alknopfler/alkcli/vpn"
	"github.com/spf13/viper"
	"time"
)

type Profile struct {
	listCmd     []string
	invocations map[string]string
}

func NewProfile(name string) *Profile {
	listInput := viper.GetStringSlice(cm.PROFILE + "." + name + "." + cm.LIST)
	inv := make(map[string]string)

	for _, val := range listInput {
		inv[val] = viper.GetString(cm.PROFILE + "." + name + "." + val)
	}
	p := &Profile{
		listCmd:     listInput,
		invocations: inv,
	}
	return p
}

func (p *Profile) ExecProfile() {
	for _, val := range p.listCmd {

		switch val {
		case cm.CONNECTION:
			fmt.Println("[alkcli] Task: Connection launched")
			conn.NewConnection(p.invocations[cm.CONNECTION]).ExecConnection()
		case cm.VPN:
			fmt.Println("[alkcli] Task: VPN launched")
			vpn.NewVpn(p.invocations[cm.VPN]).ExecVpn()
			time.Sleep(10 * time.Second)
		case cm.TUNNEL:
			fmt.Println("[alkcli] Task: Tunnel launched")
			tunnel.NewTunnel((p.invocations[cm.TUNNEL])).ExecTunnel()
		case cm.EXEC:
			fmt.Println("[alkcli] Task: Exec launched")
			exec.NewExec(p.invocations[cm.EXEC]).ExecLine()
		}
		time.Sleep(3 * time.Second) //time between commands
	}
}
