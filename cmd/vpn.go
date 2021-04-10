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
	"fmt"
	cm "github.com/alknopfler/alkcli/configMgmt"
	"github.com/alknopfler/alkcli/helper"
	v "github.com/alknopfler/alkcli/vpn"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var vpnCmd = &cobra.Command{
	Use:   "vpn",
	Short: "The 'vpn' command will create a vpn using the command specified in config file",
	Long: `The 'vpn' command will create a vpn to establish the connection to the target host parsing the config file:

- You could create a tunnel to the destination network jumping through target host:
     # alkcli vpn <provider1>
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			fmt.Println(cm.VPN + "." + args[0] + "." + cm.CMD)
			if args[0] == "" || !viper.IsSet(cm.VPN+"."+args[0]+"."+cm.CMD) {
				helper.HandleError(errors.New("VPN must have set a <command> param to select from config file. Use -h or --help"))
				return
			}

			w := v.NewVpn(args[0])
			w.ExecVpn()
		}
	},
}

func init() {
	rootCmd.AddCommand(vpnCmd)
}
