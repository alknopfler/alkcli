/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	c "github.com/alknopfler/alkcli/connect"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "The 'connect' command will connect to the target",
	Long: `The 'connect' command exec a command to establish the connection to the target host:

- You could connect to a host with the option declared in the config file just doing:
     # alkcli connect <target>
- If you want, also, you could override some params
     # alkcli connect <target> --user xxxx
`,

	Run: func(cmd *cobra.Command, args []string) {

		x, _ := cmd.Flags().GetBool("x11")
		privKey, _ := cmd.Flags().GetString("ssh-priv-key")
		user, _ := cmd.Flags().GetString("user")

		conn :=
			conn.ExecConnection(args, x, user, privKey)
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.PersistentFlags().BoolP("x11", "X", false, "Export X11 with the connection")
	connectCmd.PersistentFlags().StringP("ssh-priv-key", "k", "", "SSH Private Key to use with the connection")
	connectCmd.PersistentFlags().StringP("user", "u", "", "Username used to establish the connection")

}
