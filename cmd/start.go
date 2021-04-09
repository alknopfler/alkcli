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
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "The 'start' subcommand is to prepare an environment.",
	Long: `You could have different environment so with this command could prepare a work environment or 
maybe your home environment. You could automate some tasks to launch as a plan automatically `,

}

func init() {
	rootCmd.AddCommand(startCmd)
}
