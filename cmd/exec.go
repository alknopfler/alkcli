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
	ex "github.com/alknopfler/alkcli/exec"
	"github.com/alknopfler/alkcli/helper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "The 'exec' command will execute the command specified in config file",
	Long: `The 'exec' command will execute the command specified in config file:

- You could create a execution created in the config file just using:
     # alkcli exec <command>
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "" || !viper.IsSet(cm.EXEC+"."+args[0]) {
				helper.HandleError(errors.New("Exec must have set a <command> param to select from config file. Use -h or --help"))
				return
			}

			e := ex.NewExec(args[0])
			e.ExecVpn()
		}
	},
}

func init() {
	rootCmd.AddCommand(execCmd)
}
