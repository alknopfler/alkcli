package connect


import (
	"os"
	"os/exec"
	"syscall"
)

func ExecConnection(command string, args []string, x11 bool, user string, privKey string) {


	binary, lookErr := exec.LookPath(command)
	if lookErr != nil {
		panic(lookErr)
	}
	syscall.Exec(binary, []string{command, args[0], valueX11, "-l " + user, "-i"}, os.Environ())

}

