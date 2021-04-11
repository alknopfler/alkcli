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
	prof "github.com/alknopfler/alkcli/profile"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "The 'profile' command will exec a list of things specified in config file",
	Long: `The 'profile' command will create a list of things to setup an environment specified in the file:

- You could launch a profile to the destination network jumping through target host:
     # alkcli vpn <provider1>
`,

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			if args[0] == "" || !viper.IsSet(cm.PROFILE+"."+args[0]) {
				helper.HandleError(errors.New("profile must have set a <profile-name> param to select from config file. Use -h or --help"))
				return
			}

			p := prof.NewProfile(args[0])
			p.ExecProfile()
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
