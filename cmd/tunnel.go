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
	"errors"
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/alknopfler/alkcli/helper"
	tunn "github.com/alknopfler/alkcli/tunnel"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var tunnelCmd = &cobra.Command{
	Use:   "tunnel",
	Short: "The 'tunnel' command will create a tunnel using sshuttle to the target host",
	Long: `The 'tunnel' command will create a tunnel to establish the connection to the target host parsing the network:

- You could create a tunnel to the destination network jumping through target host:
     # alkcli tunnel <target>
- If you want, also, you could override some params
     # alkcli tunnel <target> --network <cidr>
`,

	Run: func(cmd *cobra.Command, args []string) {
		params := map[string]string{}
		if len(args) > 0 {
			network, _ := cmd.Flags().GetString("network")

			if !viper.IsSet(cm.TUNNEL + "." + cm.CMD) {
				helper.HandleError(errors.New("Tunnel must have set a <cmd> key into <tunnel> yaml section. Use -h or --help"))
				return
			}
			if args[0] == "" || !viper.IsSet(cm.TUNNEL+"."+args[0]+"."+cm.TARGET) {
				helper.HandleError(errors.New("Tunnel must have set a <target> param to select from config file. Use -h or --help"))
				return
			}

			if network != "" {
				params["network"] = network
			}

			target := viper.GetString(cm.TUNNEL + "." + args[0] + "." + cm.TARGET)

			if len(params) > 0 {
				t := tunn.NewTunnel(target, tunn.WithParams(params))
				t.ExecTunnel()
				return
			}

			t := tunn.NewTunnel(target)
			t.ExecTunnel()
		}
	},
}

func init() {

	rootCmd.AddCommand(tunnelCmd)

	tunnelCmd.PersistentFlags().StringP("network", "n", "", "Network to reach creating the tunnel with the target host")

}
