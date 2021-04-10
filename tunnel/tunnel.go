package tunnel

import (
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

type Tunnel struct {
	Command string
	Target  string
	Network string
}

type TunnelOptions func(t *Tunnel)

func NewTunnel(target string, opts ...TunnelOptions) *Tunnel {
	tunn := &Tunnel{
		Command: viper.GetString(cm.TUNNEL + "." + cm.CMD),
		Target:  viper.GetString(cm.TUNNEL + "." + target + "." + cm.TARGET),
		Network: viper.GetString(cm.TUNNEL + "." + target + "." + cm.NETWORK),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(tunn)
		}
		return tunn
	}
	return tunn
}

func WithParams(params map[string]string) TunnelOptions {
	return func(t *Tunnel) {
		t.Network = params["network"]
	}
}

func (t *Tunnel) ExecTunnel() {
	cmd := exec.Command(t.Command, "-D", "-r", t.Target, t.Network)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}
