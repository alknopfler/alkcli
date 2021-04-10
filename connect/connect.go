package connect

import (
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/alknopfler/alkcli/helper"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"syscall"
)

type Connection struct {
	Command string
	Target  string
	X11     string
	User    string
	PrivKey string
}

type ConnectionOptions func(c *Connection)

func NewConnection(opts ...ConnectionOptions) *Connection {
	conn := &Connection{
		Command: viper.GetString(cm.CONNECTION + "." + cm.CMD),
	}
	for _, opt := range opts {
		opt(conn)
	}
	return conn
}

func WithParams(params map[string]string) ConnectionOptions {
	return func(c *Connection) {
		c.Target = params["target"]
		c.X11 = params["x11"]
		c.User = params["user"]
		c.PrivKey = params["privKey"]
	}
}

func (c *Connection) ExecConnection() {

	binCmd, lookErr := exec.LookPath(c.Command)
	if lookErr != nil {
		helper.HandleError(lookErr)
		return
	}
	err := syscall.Exec(binCmd, []string{c.X11, c.Target, "-l " + c.User, "-i " + c.PrivKey}, os.Environ())
	if err != nil {
		helper.HandleError(err)
		return
	}
}
