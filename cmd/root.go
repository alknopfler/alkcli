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
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "alkcli",
	Short: "Alkcli is a simplifier tool just to automate basic daily tasks",
	Long: `With alkcli you could add options to the configuration file in order
expand the functionality to create new tasks automated .
The config file must be in the same folder than the binary but it hidden.
Example:
 - .alkcli-config-example.yml in the github repo. 
 - Just rename it to be .alkcli-config.yml with your params
 - Located in the same folder than the binary alkcli`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {

	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("ALKCLI")
	viper.AddConfigPath(".")
	viper.SetConfigName(".alkcli-config.yml")

	//helper.HandleError(viper.BindEnv("uno"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error Getting configuration file: ", viper.ConfigFileUsed())
	}
}
