package connect

import (
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/spf13/viper"
	"os"
	"os/exec"
)

type Connection struct {
	Command string
	Target  string
	X11     string
	User    string
	PrivKey string
}

type ConnectionOptions func(c *Connection)

func NewConnection(target string, opts ...ConnectionOptions) *Connection {
	conn := &Connection{
		Command: viper.GetString(cm.CONNECTION + "." + cm.CMD),
		Target:  viper.GetString(cm.CONNECTION + "." + target + "." + cm.TARGET),
		User:    viper.GetString(cm.CONNECTION + "." + target + "." + cm.USER),
		PrivKey: viper.GetString(cm.CONNECTION + "." + target + "." + cm.KEY),
		X11:     viper.GetString(cm.CONNECTION + "." + target + "." + cm.X),
	}
	if len(opts) > 0 {
		for _, opt := range opts {
			opt(conn)
		}
		return conn
	}
	return conn
}

func WithParams(params map[string]string) ConnectionOptions {
	return func(c *Connection) {
		c.X11 = params["x11"]
		c.User = params["user"]
		c.PrivKey = params["privKey"]
	}
}

func (c *Connection) ExecConnection() {
	cmd := exec.Command(c.Command, translateX11(c.X11)+translatePrivKey(c.PrivKey)+c.User+"@"+c.Target)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func translateX11(s string) string {
	if s == "true" {
		return " -X "
	} else {
		return ""
	}
}
func translatePrivKey(s string) string {
	if s != "" {
		return " -i" + s + " "
	} else {
		return ""
	}
}
