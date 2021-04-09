package connect

import (
	"errors"
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/alknopfler/alkcli/helper"
	"github.com/spf13/viper"
)

type Connection struct {
	Command string
	Target  string
	X11     bool
	User    string
	PrivKey string
}

type ConnectionOptions func(c *Connection)

func NewConnection(opts ...ConnectionOptions) *Connection {
	conn := &Connection{}

}

func (c *Connection) ExecConnection() {

	// alkcli connect connection.cmd args[0] x11

	if !viper.IsSet(cm.CONNECTION + "." + cm.CMD) {
		helper.HandleError(errors.New("Connection must have set a <cmd> key into <connection> yaml section"))
		return
	}
	if c.Target == "" || !viper.IsSet(cm.CONNECTION+"."+c.Target+"."+cm.TARGET) {
		helper.HandleError(errors.New("Connection must have set a <target> key into <connection> yaml section or maybe the command is not complete"))
		return
	}
	if c.User != "" {
		u := c.User
	} else {
		u := viper.GetString(cm.CONNECTION + "." + cm.CMD)
	}

	command := viper.GetString(cm.CONNECTION + "." + cm.CMD)
	target := viper.GetString(cm.CONNECTION + "." + args[0] + "." + cm.TARGET)
	x := x11 || viper.GetBool(cm.CONNECTION+"."+args[0]+"."+cm.X)

	/*binary, lookErr := exec.LookPath(command)
	if lookErr != nil {
		panic(lookErr)
	}
	syscall.Exec(binary, []string{command, args[0], valueX11, "-l " + user, "-i"}, os.Environ())
	*/
}
